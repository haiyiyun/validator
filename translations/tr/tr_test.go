package tr

import (
	"testing"
	"time"

	"github.com/haiyiyun/validator"
	. "github.com/haiyiyun/validator/assert"
	turkish "github.com/haiyiyun/validator/locales/tr"
	ut "github.com/haiyiyun/validator/universal-translator"
)

func TestTranslations(t *testing.T) {

	tr := turkish.New()
	uni := ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")

	validate := validator.New()

	err := RegisterDefaultTranslations(validate, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string            `validate:"required"`
		RequiredNumber    int               `validate:"required"`
		RequiredMultiple  []string          `validate:"required"`
		LenString         string            `validate:"len=1"`
		LenNumber         float64           `validate:"len=1113.00"`
		LenMultiple       []string          `validate:"len=7"`
		MinString         string            `validate:"min=1"`
		MinNumber         float64           `validate:"min=1113.00"`
		MinMultiple       []string          `validate:"min=7"`
		MaxString         string            `validate:"max=3"`
		MaxNumber         float64           `validate:"max=1113.00"`
		MaxMultiple       []string          `validate:"max=7"`
		EqString          string            `validate:"eq=3"`
		EqNumber          float64           `validate:"eq=2.33"`
		EqMultiple        []string          `validate:"eq=7"`
		NeString          string            `validate:"ne="`
		NeNumber          float64           `validate:"ne=0.00"`
		NeMultiple        []string          `validate:"ne=0"`
		LtString          string            `validate:"lt=3"`
		LtNumber          float64           `validate:"lt=5.56"`
		LtMultiple        []string          `validate:"lt=2"`
		LtTime            time.Time         `validate:"lt"`
		LteString         string            `validate:"lte=3"`
		LteNumber         float64           `validate:"lte=5.56"`
		LteMultiple       []string          `validate:"lte=2"`
		LteTime           time.Time         `validate:"lte"`
		GtString          string            `validate:"gt=3"`
		GtNumber          float64           `validate:"gt=5.56"`
		GtMultiple        []string          `validate:"gt=2"`
		GtTime            time.Time         `validate:"gt"`
		GteString         string            `validate:"gte=3"`
		GteNumber         float64           `validate:"gte=5.56"`
		GteMultiple       []string          `validate:"gte=2"`
		GteTime           time.Time         `validate:"gte"`
		EqFieldString     string            `validate:"eqfield=MaxString"`
		EqCSFieldString   string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string            `validate:"nefield=EqFieldString"`
		GtFieldString     string            `validate:"gtfield=MaxString"`
		GteFieldString    string            `validate:"gtefield=MaxString"`
		LtFieldString     string            `validate:"ltfield=MaxString"`
		LteFieldString    string            `validate:"ltefield=MaxString"`
		AlphaString       string            `validate:"alpha"`
		AlphanumString    string            `validate:"alphanum"`
		NumericString     string            `validate:"numeric"`
		NumberString      string            `validate:"number"`
		HexadecimalString string            `validate:"hexadecimal"`
		HexColorString    string            `validate:"hexcolor"`
		RGBColorString    string            `validate:"rgb"`
		RGBAColorString   string            `validate:"rgba"`
		HSLColorString    string            `validate:"hsl"`
		HSLAColorString   string            `validate:"hsla"`
		Email             string            `validate:"email"`
		URL               string            `validate:"url"`
		URI               string            `validate:"uri"`
		Base64            string            `validate:"base64"`
		Contains          string            `validate:"contains=purpose"`
		ContainsAny       string            `validate:"containsany=!@#$"`
		Excludes          string            `validate:"excludes=text"`
		ExcludesAll       string            `validate:"excludesall=!@#$"`
		ExcludesRune      string            `validate:"excludesrune=???"`
		ISBN              string            `validate:"isbn"`
		ISBN10            string            `validate:"isbn10"`
		ISBN13            string            `validate:"isbn13"`
		UUID              string            `validate:"uuid"`
		UUID3             string            `validate:"uuid3"`
		UUID4             string            `validate:"uuid4"`
		UUID5             string            `validate:"uuid5"`
		ASCII             string            `validate:"ascii"`
		PrintableASCII    string            `validate:"printascii"`
		MultiByte         string            `validate:"multibyte"`
		DataURI           string            `validate:"datauri"`
		Latitude          string            `validate:"latitude"`
		Longitude         string            `validate:"longitude"`
		SSN               string            `validate:"ssn"`
		IP                string            `validate:"ip"`
		IPv4              string            `validate:"ipv4"`
		IPv6              string            `validate:"ipv6"`
		CIDR              string            `validate:"cidr"`
		CIDRv4            string            `validate:"cidrv4"`
		CIDRv6            string            `validate:"cidrv6"`
		TCPAddr           string            `validate:"tcp_addr"`
		TCPAddrv4         string            `validate:"tcp4_addr"`
		TCPAddrv6         string            `validate:"tcp6_addr"`
		UDPAddr           string            `validate:"udp_addr"`
		UDPAddrv4         string            `validate:"udp4_addr"`
		UDPAddrv6         string            `validate:"udp6_addr"`
		IPAddr            string            `validate:"ip_addr"`
		IPAddrv4          string            `validate:"ip4_addr"`
		IPAddrv6          string            `validate:"ip6_addr"`
		UinxAddr          string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string            `validate:"mac"`
		IsColor           string            `validate:"iscolor"`
		StrPtrMinLen      *string           `validate:"min=10"`
		StrPtrMaxLen      *string           `validate:"max=1"`
		StrPtrLen         *string           `validate:"len=2"`
		StrPtrLt          *string           `validate:"lt=1"`
		StrPtrLte         *string           `validate:"lte=1"`
		StrPtrGt          *string           `validate:"gt=10"`
		StrPtrGte         *string           `validate:"gte=10"`
		OneOfString       string            `validate:"oneof=red green"`
		OneOfInt          int               `validate:"oneof=5 63"`
		UniqueSlice       []string          `validate:"unique"`
		UniqueArray       [3]string         `validate:"unique"`
		UniqueMap         map[string]string `validate:"unique"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)

	test.LtString = "1234"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteTime = time.Now().Add(time.Hour * 24)

	test.LtFieldString = "12345"
	test.LteFieldString = "12345"

	test.LtCSFieldString = "1234"
	test.LteCSFieldString = "1234"

	test.AlphaString = "abc3"
	test.AlphanumString = "abc3!"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ???"

	test.ASCII = "????????????"
	test.PrintableASCII = "????????????"

	test.MultiByte = "1234feerf"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(validator.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.IsColor",
			expected: "IsColor ge??erli bir renk olmal??d??r",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC ge??erli bir MAC adresi i??ermelidir",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr ????z??lebilir bir IP adresi olmal??d??r",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 ????z??lebilir bir IPv4 adresi olmal??d??r",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 ????z??lebilir bir IPv6 adresi olmal??d??r",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr ge??erli bir UDP adresi olmal??d??r",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 ge??erli bir IPv4 UDP adresi olmal??d??r",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 ge??erli bir IPv6 UDP adresi olmal??d??r",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr ge??erli bir TCP adresi olmal??d??r",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 ge??erli bir IPv4 TCP adresi olmal??d??r",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 ge??erli bir IPv6 TCP adresi olmal??d??r",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR ge??erli bir CIDR g??sterimi i??ermelidir",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 bir IPv4 adresi i??in ge??erli bir CIDR g??sterimi i??ermelidir",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 bir IPv6 adresi i??in ge??erli bir CIDR g??sterimi i??ermelidir",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN ge??erli bir SSN numaras?? olmal??d??r",
		},
		{
			ns:       "Test.IP",
			expected: "IP ge??erli bir IP adresi olmal??d??r",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 ge??erli bir IPv4 adresi olmal??d??r",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 ge??erli bir IPv6 adresi olmal??d??r",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI ge??erli bir Veri URI i??ermelidir",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude ge??erli bir enlem koordinat?? i??ermelidir",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude ge??erli bir boylam koordinat?? i??ermelidir",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte ??ok baytl?? karakterler i??ermelidir",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII yaln??zca ascii karakterler i??ermelidir",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII yaln??zca yazd??r??labilir ascii karakterleri i??ermelidir",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID ge??erli bir UUID olmal??d??r",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 ge??erli bir s??r??m 3 UUID olmal??d??r",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 ge??erli bir s??r??m 4 UUID olmal??d??r",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 ge??erli bir s??r??m 5 UUID olmal??d??r",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN ge??erli bir ISBN numaras?? olmal??d??r",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 ge??erli bir ISBN-10 numaras?? olmal??d??r",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 ge??erli bir ISBN-13 numaras?? olmal??d??r",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes, 'text' metnini i??eremez",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll, '!@#$' karakterlerinden hi??birini i??eremez",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune, '???' ifadesini i??eremez",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny, '!@#$' karakterlerinden en az birini i??ermelidir",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains, 'purpose' metnini i??ermelidir",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 ge??erli bir Base64 karakter dizesi olmal??d??r",
		},
		{
			ns:       "Test.Email",
			expected: "Email ge??erli bir e-posta adresi olmal??d??r",
		},
		{
			ns:       "Test.URL",
			expected: "URL ge??erli bir URL olmal??d??r",
		},
		{
			ns:       "Test.URI",
			expected: "URI ge??erli bir URI olmal??d??r",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString ge??erli bir RGB rengi olmal??d??r",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString ge??erli bir RGBA rengi olmal??d??r",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString ge??erli bir HSL rengi olmal??d??r",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString ge??erli bir HSLA rengi olmal??d??r",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString ge??erli bir onalt??l??k olmal??d??r",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString ge??erli bir HEX rengi olmal??d??r",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString ge??erli bir say?? olmal??d??r",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString ge??erli bir say??sal de??er olmal??d??r",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString yaln??zca alfan??merik karakterler i??erebilir",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString yaln??zca alfabetik karakterler i??erebilir",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString, MaxString de??erinden k??????k olmal??d??r",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString, MaxString de??erinden k??????k veya ona e??it olmal??d??r",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString, MaxString de??erinden b??y??k olmal??d??r",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString, MaxString de??erinden b??y??k veya ona e??it olmal??d??r",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString, EqFieldString de??erine e??it olmamal??d??r",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString, Inner.LtCSFieldString de??erinden k??????k olmal??d??r",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString, Inner.LteCSFieldString de??erinden k??????k veya ona e??it olmal??d??r",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString, Inner.GtCSFieldString de??erinden b??y??k olmal??d??r",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString, Inner.GteCSFieldString de??erinden k??????k veya ona e??it olmal??d??r",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString, Inner.NeCSFieldString de??erine e??it olmamal??d??r",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString, Inner.EqCSFieldString de??erine e??it olmal??d??r",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString, MaxString de??erine e??it olmal??d??r",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString en az 3 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber, 5,56 veya daha b??y??k olmal??d??r",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple en az 2 ????e i??ermelidir",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime ge??erli Tarih ve Saatten b??y??k veya ona e??it olmal??d??r",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString, 3 karakter uzunlu??undan fazla olmal??d??r",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber, 5,56 de??erinden b??y??k olmal??d??r",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple, 2 ????eden daha fazla i??ermelidir",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime ge??erli Tarih ve Saatten b??y??k olmal??d??r",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString en fazla 3 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber, 5,56 veya daha az olmal??d??r",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple, maksimum 2 ????e i??ermelidir",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime ge??erli Tarih ve Saate e??it veya daha k??????k olmal??d??r",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString, 3 karakter uzunlu??undan daha az olmal??d??r",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber, 5,56 de??erinden k??????k olmal??d??r",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple, 2 ????eden daha az i??ermelidir",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime ge??erli Tarih ve Saatten daha az olmal??d??r",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString,  de??erine e??it olmamal??d??r",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber, 0.00 de??erine e??it olmamal??d??r",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple, 0 de??erine e??it olmamal??d??r",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString, 3 de??erine e??it de??il",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber, 2.33 de??erine e??it de??il",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple, 7 de??erine e??it de??il",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString uzunlu??u en fazla 3 karakter olmal??d??r",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber, 1.113,00 veya daha az olmal??d??r",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple maksimum 7 ????e i??ermelidir",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString en az 1 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber, 1.113,00 veya daha b??y??k olmal??d??r",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple en az 7 ????e i??ermelidir",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString uzunlu??u 1 karakter olmal??d??r",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber, 1.113,00 de??erine e??it olmal??d??r",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple, 7 ????e i??ermelidir",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString zorunlu bir aland??r",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber zorunlu bir aland??r",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple zorunlu bir aland??r",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen en az 10 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen uzunlu??u en fazla 1 karakter olmal??d??r",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen uzunlu??u 2 karakter olmal??d??r",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt, 1 karakter uzunlu??undan daha az olmal??d??r",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte en fazla 1 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt, 10 karakter uzunlu??undan fazla olmal??d??r",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte en az 10 karakter uzunlu??unda olmal??d??r",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString, [red green]'dan biri olmal??d??r",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt, [5 63]'dan biri olmal??d??r",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice benzersiz de??erler i??ermelidir",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray benzersiz de??erler i??ermelidir",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap benzersiz de??erler i??ermelidir",
		},
	}

	for _, tt := range tests {

		var fe validator.FieldError

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}

}
