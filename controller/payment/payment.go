package payment

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/model"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/proxy"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/controller/booking"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Body struct {
	Amount          string                   `json:"amount"`
	BookingProducts []booking.BookingProduct `json:"bookingProducts"`
}

var omisePublicKey string = os.Getenv("OMISE_PUBLIC_KEY")
var omiseSecretKey string = os.Getenv("OMISE_SECRET_KEY")

func PaymentQrCode(c echo.Context, db *gorm.DB) error {

	var body Body = Body{}
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// ทำ HTTP POST
	url := fmt.Sprintf("%s/charges", proxy.OMISE_URL)
	bodyPayload := fmt.Sprintf("amount=%s&currency=THB&source[type]=promptpay", body.Amount)
	payload := strings.NewReader(bodyPayload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating request"})
	}

	// ใส่ Authorization header
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(omiseSecretKey+":")))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// ทำ HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error making request"})
	}
	defer resp.Body.Close()

	// อ่าน response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	omiseCharge, err := model.UnmarshalOmiseCharge([]byte(responseBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error parsing response"})
	}

	const layout = "2006-01-02T15:04:05Z"
	convertCreateAt, err := time.Parse(layout, omiseCharge.Source.ScannableCode.Image.CreatedAt)

	var payment Payment = Payment{
		PaymentID:       omiseCharge.ID,
		Amount:          int(omiseCharge.Amount),
		BookingProducts: body.BookingProducts,
		Source:          omiseCharge.Source.ScannableCode.Image.DownloadURI,
		PaymentMethod:   "promptpay",
		Status:          omiseCharge.Status,
		CreatedAt:       convertCreateAt,
		UpdatedAt:       time.Now(),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    payment,
	})

}

func GetPaymentStatus(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	url := fmt.Sprintf("%s/charges/%s", proxy.OMISE_URL, id)

	req, err := http.NewRequest("GET", url, strings.NewReader(""))

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(omiseSecretKey+":")))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating request"})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error making request"})
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	// if status success save booking to database else reject

	omiseCharge, err := model.UnmarshalOmiseCharge([]byte(responseBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error parsing response"})
	}

	const layout = "2006-01-02T15:04:05Z"
	convertCreateAt, err := time.Parse(layout, omiseCharge.Source.ScannableCode.Image.CreatedAt)

	var payment Payment = Payment{
		PaymentID:       omiseCharge.ID,
		Amount:          int(omiseCharge.Amount),
		BookingProducts: []booking.BookingProduct{},
		Source:          omiseCharge.Source.ScannableCode.Image.DownloadURI,
		PaymentMethod:   "promptpay",
		Status:          omiseCharge.Status,
		CreatedAt:       convertCreateAt,
		UpdatedAt:       time.Now(),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    payment,
	})
}
