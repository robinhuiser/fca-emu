package generator

import (
	"bytes"
	"embed"
	"encoding/base64"
	"fmt"
	"math"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/divan/num2words"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

//go:embed imgs/bank-of-pi-logo.png
var logo embed.FS

func renderCheck(f *gofakeit.Faker, fullname string, street string, city string, pc string, amount float64, designated string, memo string) (bytes.Buffer, error) {
	m := pdf.NewMaroto(consts.Landscape, consts.A3)
	l, _ := logo.ReadFile("imgs/bank-of-pi-logo.png")

	// Title row
	m.Row(60, func() {
		m.Col(2, func() {
			_ = m.Base64Image(base64.StdEncoding.EncodeToString(l), consts.Png, props.Rect{
				Center:  true,
				Percent: 90,
			})
		})
		m.Col(8, func() {
			m.Text("Bank of Pi, Inc.", props.Text{
				Top:         8,
				Size:        60,
				Extrapolate: true,
			})
			m.Text("1000 Shipping Gopher Golang TN 3691234 GopherLand (GL)", props.Text{
				Size: 24,
				Top:  34,
			})
		})
		m.ColSpace(4)
	})

	m.Line(10)

	// Name of check owner, date of writing check
	m.Row(30, func() {
		m.Col(4, func() {
			m.Text(fullname, props.Text{
				Size:  15,
				Top:   4,
				Style: consts.Bold,
			})
			m.Text(street, props.Text{
				Size: 15,
				Top:  10,
			})
			m.Text(city+" "+pc, props.Text{
				Size: 15,
				Top:  16,
			})

		})
		m.ColSpace(3)
		m.Col(2, func() {
			m.Text("DATE", props.Text{
				Size:  15,
				Top:   6,
				Align: consts.Right,
				Style: consts.Bold,
			})
		})

		m.Col(2, func() {
			m.Text(" 07/01/2018", props.Text{
				Size:   25,
				Top:    6,
				Family: consts.Courier,
				Align:  consts.Right,
			})
			m.Text("   ______________", props.Text{
				Size:  20,
				Top:   12,
				Align: consts.Right,
			})
		})
	})

	// Pay to + amount
	m.Row(30, func() {
		m.Col(1, func() {
			m.Text("PAY TO THE ORDER OF", props.Text{
				Size:  15,
				Top:   4,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.ColSpace(1)
		m.Col(4, func() {
			m.Text(designated, props.Text{
				Size:   25,
				Top:    4,
				Family: consts.Courier,
			})
			m.Text("_____________________________________________________________________", props.Text{
				Size: 15,
				Top:  12,
			})
		})
		m.ColSpace(2)
		m.Col(1, func() {
			m.Text("$", props.Text{
				Size:  25,
				Top:   4,
				Align: consts.Right,
			})
		})
		m.Col(2, func() {
			m.Text(fmt.Sprintf("  %.2f", amount), props.Text{
				Size:   25,
				Top:    4,
				Family: consts.Courier,
				Align:  consts.Right,
			})
			m.Text("____________________", props.Text{
				Size:  15,
				Top:   12,
				Align: consts.Right,
			})
		})
	})

	// Amount in words
	m.Row(40, func() {
		m.ColSpace(1)
		ipart, decpart := math.Modf(amount)
		c := 0

		cnts := "cents"

		if decpart > 0 {
			c, _ = strconv.Atoi(fmt.Sprintf("%.2f", amount-float64(ipart))[2:])
		}

		if c == 1 || c == 0 {
			cnts = "cent"
		}

		m.Col(10, func() {
			m.Text(num2words.Convert(int(ipart))+" dollars and "+num2words.Convert(c)+" "+cnts, props.Text{
				Size:            25,
				Top:             4,
				Family:          consts.Courier,
				Align:           consts.Left,
				VerticalPadding: 6,
			})
			m.Text("________________________________________________________________________________________________________________", props.Text{
				Size:  15,
				Top:   12,
				Align: consts.Left,
			})
			m.Text("________________________________________________________________________________________________________________", props.Text{
				Size:  15,
				Top:   26,
				Align: consts.Left,
			})
		})
	})

	// Memo and signature
	m.Row(30, func() {
		m.Col(1, func() {
			m.Text("MEMO", props.Text{
				Size:  15,
				Top:   4,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(4, func() {
			m.Text(memo, props.Text{
				Size:   25,
				Top:    4,
				Family: consts.Courier,
			})
			m.Text("____________________________________________", props.Text{
				Size: 15,
				Top:  12,
			})
		})
		m.Col(2, func() {
			m.Text("SIGNATURE", props.Text{
				Size:  15,
				Top:   4,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(4, func() {
			m.Text("", props.Text{
				Size:   25,
				Top:    4,
				Family: consts.Courier,
			})
			m.Text("____________________________________________", props.Text{
				Size: 15,
				Top:  12,
			})
		})
	})

	m.Line(10)

	// Barcode
	m.Row(50, func() {
		m.Col(12, func() {
			code := strconv.Itoa(f.Number(100000000, 999999999)) +
				"\t\t\t" + strconv.Itoa(f.Number(100000000, 999999999)) +
				"\t\t\t" + strconv.Itoa(f.Number(1000, 9999))
			_ = m.Barcode(code, props.Barcode{
				Center:  true,
				Percent: 40,
			})
			m.Text(code, props.Text{
				Size:   17,
				Align:  consts.Center,
				Family: consts.Courier,
				Top:    35,
			})
		})
	})

	r, err := m.Output()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("failed to serialize: %w", err)
	}
	return r, nil
}
