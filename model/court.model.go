package model

import "time"

type Court struct {
	ID          int       `json:"courts_id"`          // หรือ json:"courts_id" ถ้าคุณต้องการใช้กับ JSON
	Name        string    `json:"courts_name"`        // ใช้ tags เหล่านี้ในการ map กับชื่อคอลัมน์ใน SQL
	Description string    `json:"courts_description"` // สามารถเปลี่ยน `json` เป็น `json` หรืออื่นๆ ตามที่คุณต้องการ
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
