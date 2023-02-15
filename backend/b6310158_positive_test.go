package backend

import (
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
	"testing"
)

type Customer struct {
	gorm.Model
	Name       string //`valid:"required~ผู้ใช้ไม้ได้กรอกข้อมูลชื่อ"`
	Email      string
	CustomerID string //`valid:"matches(L\\b{7}$)~กรอก CustomerID ขึ้นต้นด้วย L หรือ M หรือ H ตามด้วยเลข 7 หลัก"`
}

// Name ไม่เป็นค่าว่าง
func testName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	a := Customer{
		Name:       "",
		Email:      "Natthapong@gmail.com",
		CustomerID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(a)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("ผู้ใช้ไม้ได้กรอกข้อมูลชื่อ"))
}

// CustomerID ขึ้นต้นด้วย L หรือ M หรือ H ตามด้วยเลข 7 หลัก
// func testCutomerID(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 	a := Customer{
// 		Name:       "Natthapong",
// 		Email:      "Natthapong@gmail.com",
// 		CustomerID: "1234567",
// 	}

// 	ok, err := govalidator.ValidateStruct(a)
// 	g.Expect(ok).NotTo(gomega.BeTrue())
// 	g.Expect(err).ToNot(gomega.BeNil())
// 	g.Expect(err.Error()).To(gomega.Equal("กรอก CustomerID ขึ้นต้นด้วย L หรือ M หรือ H ตามด้วยเลข 7 หลัก"))
// }

// สามารถบันทึกข้อมูลได้
// func testAll(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 	a := Customer{
// 		Name:       "Natthapong",
// 		Email:      "Natthapong@gmail.com",
// 		CustomerID: "L1234567",
// 	}

// 	ok, err := govalidator.ValidateStruct(a)
// 	g.Expect(ok).To(gomega.BeTrue())
// 	g.Expect(err).To(gomega.BeNil())
// }
