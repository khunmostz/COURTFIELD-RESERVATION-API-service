package payment

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/artifact/proxy"
	"github.com/khunmostz24/COURTFIELD-RESERVATION-API-service/model"
	"github.com/labstack/echo/v4"
)

type Body struct {
	Amount      string    `json:"amount"`
	Reservation time.Time `json:"reservation_time"`
}

var omisePublicKey string = os.Getenv("OMISE_PUBLIC_KEY")
var omiseSecretKey string = os.Getenv("OMISE_SECRET_KEY")

func PaymentQrCode(c echo.Context, db *sql.DB) error {

	var body Body = Body{}
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// ทำ HTTP POST
	url := fmt.Sprintf("%s/charges", proxy.OMISE_URL)
	bodyPayload := fmt.Sprintf("amount=%s&currency=THB&source[type]=promptpay", body.Amount)
	fmt.Printf(bodyPayload)
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    omiseCharge,
	})

}

func GetPaymentStatus(c echo.Context, db *sql.DB) error {
	id := c.QueryParam("id")
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    omiseCharge,
	})
}
