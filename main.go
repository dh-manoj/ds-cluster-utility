package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const startPortNumber = 54000

var countryCodes = []string{"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MK", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW", "STG", "TEST"}
var portNumbers = map[string]int{}

func init() {
	portNumbers = make(map[string]int)

	//initialize with port numbers
	for i, code := range countryCodes {
		portNumbers[code] = startPortNumber + i
	}
}

var countryName = map[string]string{
	"AF":   "Afghanistan",
	"AX":   "Åland Islands",
	"AL":   "Albania",
	"DZ":   "Algeria",
	"AS":   "American Samoa",
	"AD":   "Andorra",
	"AO":   "Angola",
	"AI":   "Anguilla",
	"AQ":   "Antarctica",
	"AG":   "Antigua and Barbuda",
	"AR":   "Argentina",
	"AM":   "Armenia",
	"AW":   "Aruba",
	"AU":   "Australia",
	"AT":   "Austria",
	"AZ":   "Azerbaijan",
	"BS":   "Bahamas",
	"BH":   "Bahrain",
	"BD":   "Bangladesh",
	"BB":   "Barbados",
	"BY":   "Belarus",
	"BE":   "Belgium",
	"BZ":   "Belize",
	"BJ":   "Benin",
	"BM":   "Bermuda",
	"BT":   "Bhutan",
	"BO":   "Bolivia (Plurinational State of)",
	"BQ":   "Bonaire, Sint Eustatius and Saba",
	"BA":   "Bosnia and Herzegovina",
	"BW":   "Botswana",
	"BV":   "Bouvet Island",
	"BR":   "Brazil",
	"IO":   "British Indian Ocean Territory",
	"BN":   "Brunei Darussalam",
	"BG":   "Bulgaria",
	"BF":   "Burkina Faso",
	"BI":   "Burundi",
	"CV":   "Cabo Verde",
	"KH":   "Cambodia",
	"CM":   "Cameroon",
	"CA":   "Canada",
	"KY":   "Cayman Islands",
	"CF":   "Central African Republic",
	"TD":   "Chad",
	"CL":   "Chile",
	"CN":   "China",
	"CX":   "Christmas Island",
	"CC":   "Cocos (Keeling) Islands",
	"CO":   "Colombia",
	"KM":   "Comoros",
	"CG":   "Congo",
	"CD":   "Congo, Democratic Republic of the",
	"CK":   "Cook Islands",
	"CR":   "Costa Rica",
	"CI":   "Côte d'Ivoire",
	"HR":   "Croatia",
	"CU":   "Cuba",
	"CW":   "Curaçao",
	"CY":   "Cyprus",
	"CZ":   "Czechia",
	"DK":   "Denmark",
	"DJ":   "Djibouti",
	"DM":   "Dominica",
	"DO":   "Dominican Republic",
	"EC":   "Ecuador",
	"EG":   "Egypt",
	"SV":   "El Salvador",
	"GQ":   "Equatorial Guinea",
	"ER":   "Eritrea",
	"EE":   "Estonia",
	"SZ":   "Eswatini",
	"ET":   "Ethiopia",
	"FK":   "Falkland Islands (Malvinas)",
	"FO":   "Faroe Islands",
	"FJ":   "Fiji",
	"FI":   "Finland",
	"FR":   "France",
	"GF":   "French Guiana",
	"PF":   "French Polynesia",
	"TF":   "French Southern Territories",
	"GA":   "Gabon",
	"GM":   "Gambia",
	"GE":   "Georgia",
	"DE":   "Germany",
	"GH":   "Ghana",
	"GI":   "Gibraltar",
	"GR":   "Greece",
	"GL":   "Greenland",
	"GD":   "Grenada",
	"GP":   "Guadeloupe",
	"GU":   "Guam",
	"GT":   "Guatemala",
	"GG":   "Guernsey",
	"GN":   "Guinea",
	"GW":   "Guinea-Bissau",
	"GY":   "Guyana",
	"HT":   "Haiti",
	"HM":   "Heard Island and McDonald Islands",
	"VA":   "Holy See",
	"HN":   "Honduras",
	"HK":   "Hong Kong",
	"HU":   "Hungary",
	"IS":   "Iceland",
	"IN":   "India",
	"ID":   "Indonesia",
	"IR":   "Iran (Islamic Republic of)",
	"IQ":   "Iraq",
	"IE":   "Ireland",
	"IM":   "Isle of Man",
	"IL":   "Israel",
	"IT":   "Italy",
	"JM":   "Jamaica",
	"JP":   "Japan",
	"JE":   "Jersey",
	"JO":   "Jordan",
	"KZ":   "Kazakhstan",
	"KE":   "Kenya",
	"KI":   "Kiribati",
	"KP":   "Korea (Democratic People's Republic of)",
	"KR":   "Korea, Republic of",
	"KW":   "Kuwait",
	"KG":   "Kyrgyzstan",
	"LA":   "Lao People's Democratic Republic",
	"LV":   "Latvia",
	"LB":   "Lebanon",
	"LS":   "Lesotho",
	"LR":   "Liberia",
	"LY":   "Libya",
	"LI":   "Liechtenstein",
	"LT":   "Lithuania",
	"LU":   "Luxembourg",
	"MO":   "Macao",
	"MG":   "Madagascar",
	"MW":   "Malawi",
	"MY":   "Malaysia",
	"MV":   "Maldives",
	"ML":   "Mali",
	"MT":   "Malta",
	"MH":   "Marshall Islands",
	"MQ":   "Martinique",
	"MR":   "Mauritania",
	"MU":   "Mauritius",
	"YT":   "Mayotte",
	"MX":   "Mexico",
	"FM":   "Micronesia (Federated States of)",
	"MD":   "Moldova, Republic of",
	"MC":   "Monaco",
	"MN":   "Mongolia",
	"ME":   "Montenegro",
	"MS":   "Montserrat",
	"MA":   "Morocco",
	"MZ":   "Mozambique",
	"MM":   "Myanmar",
	"NA":   "Namibia",
	"NR":   "Nauru",
	"NP":   "Nepal",
	"NL":   "Netherlands",
	"NC":   "New Caledonia",
	"NZ":   "New Zealand",
	"NI":   "Nicaragua",
	"NE":   "Niger",
	"NG":   "Nigeria",
	"NU":   "Niue",
	"NF":   "Norfolk Island",
	"MK":   "North Macedonia",
	"MP":   "Northern Mariana Islands",
	"NO":   "Norway",
	"OM":   "Oman",
	"PK":   "Pakistan",
	"PW":   "Palau",
	"PS":   "Palestine, State of",
	"PA":   "Panama",
	"PG":   "Papua New Guinea",
	"PY":   "Paraguay",
	"PE":   "Peru",
	"PH":   "Philippines",
	"PN":   "Pitcairn",
	"PL":   "Poland",
	"PT":   "Portugal",
	"PR":   "Puerto Rico",
	"QA":   "Qatar",
	"RE":   "Réunion",
	"RO":   "Romania",
	"RU":   "Russian Federation",
	"RW":   "Rwanda",
	"BL":   "Saint Barthélemy",
	"SH":   "Saint Helena, Ascension and Tristan da Cunha",
	"KN":   "Saint Kitts and Nevis",
	"LC":   "Saint Lucia",
	"MF":   "Saint Martin (French part)",
	"PM":   "Saint Pierre and Miquelon",
	"VC":   "Saint Vincent and the Grenadines",
	"WS":   "Samoa",
	"SM":   "San Marino",
	"ST":   "Sao Tome and Principe",
	"SA":   "Saudi Arabia",
	"SN":   "Senegal",
	"RS":   "Serbia",
	"SC":   "Seychelles",
	"SL":   "Sierra Leone",
	"SG":   "Singapore",
	"SX":   "Sint Maarten (Dutch part)",
	"SK":   "Slovakia",
	"SI":   "Slovenia",
	"SB":   "Solomon Islands",
	"SO":   "Somalia",
	"ZA":   "South Africa",
	"GS":   "South Georgia and the South Sandwich Islands",
	"SS":   "South Sudan",
	"ES":   "Spain",
	"LK":   "Sri Lanka",
	"SD":   "Sudan",
	"SR":   "Suriname",
	"SJ":   "Svalbard and Jan Mayen",
	"SE":   "Sweden",
	"CH":   "Switzerland",
	"SY":   "Syrian Arab Republic",
	"TW":   "Taiwan, Province of China",
	"TJ":   "Tajikistan",
	"TZ":   "Tanzania, United Republic of",
	"TH":   "Thailand",
	"TL":   "Timor-Leste",
	"TG":   "Togo",
	"TK":   "Tokelau",
	"TO":   "Tonga",
	"TT":   "Trinidad and Tobago",
	"TN":   "Tunisia",
	"TR":   "Turkey",
	"TM":   "Turkmenistan",
	"TC":   "Turks and Caicos Islands",
	"TV":   "Tuvalu",
	"UG":   "Uganda",
	"UA":   "Ukraine",
	"AE":   "United Arab Emirates",
	"GB":   "United Kingdom of Great Britain and Northern Ireland",
	"US":   "United States of America",
	"UM":   "United States Minor Outlying Islands",
	"UY":   "Uruguay",
	"UZ":   "Uzbekistan",
	"VU":   "Vanuatu",
	"VE":   "Venezuela (Bolivarian Republic of)",
	"VN":   "Viet Nam",
	"VG":   "Virgin Islands (British)",
	"VI":   "Virgin Islands (U.S.)",
	"WF":   "Wallis and Futuna",
	"EH":   "Western Sahara",
	"YE":   "Yemen",
	"ZM":   "Zambia",
	"ZW":   "Zimbabwe",
	"STG":  "Staging",
	"TEST": "Test",
}

func infraPath() string {
	path := os.Getenv("DH_DARKSTORE_INFRA_PATH")

	if path == "" {
		log.Fatalf("env DH_DARKSTORE_INFRA_PATH not set. Please set the enviroment to infra path")
	}
	return path
}

type countryInfo struct {
	code        string
	countryCode string
	region      string
	zone        string
	clusterName string
	filePath    string
}

var re = regexp.MustCompile("\"(.*?)\"")

func extractValueWithinDoubleQuotes(value string) string {
	match := re.FindStringSubmatch(value)
	return match[1]
}

func extractInfraInfo(env, filePath string) (string, string, string, string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rxCode := regexp.MustCompile("(country_code|country)[ =]*")
	rxRegion := regexp.MustCompile(fmt.Sprintf("^ *[a-z_]*region[ =]*"))
	rxZone := regexp.MustCompile(fmt.Sprintf("^ *[a-z_]*zone[ =]*"))
	rxClusterName := regexp.MustCompile("cluster_name[ =]*")

	isCodeFound := false
	isRegionFound := false
	isZoneFound := false
	isClusterNameFound := false

	code := ""
	region := ""
	zone := ""
	clusterName := ""
	version := 1
	for scanner.Scan() {
		if isCodeFound && isRegionFound && isZoneFound && isClusterNameFound {
			break
		} else if !isCodeFound && rxCode.MatchString(scanner.Text()) {
			code = extractValueWithinDoubleQuotes(scanner.Text())
			isCodeFound = true
		} else if !isRegionFound && rxRegion.MatchString(scanner.Text()) {
			region = extractValueWithinDoubleQuotes(scanner.Text())
			isRegionFound = true
		} else if !isZoneFound && rxZone.MatchString(scanner.Text()) {
			zone = extractValueWithinDoubleQuotes(scanner.Text())
			//hacky code - because of different zone name confusion in cluster-xx.tf file
			if zone == "us-east4-b" && strings.Contains(filePath, "cluster-arg") {
				zone = "us-east4-c" //argentina
			} else if zone == "europe-west1-d" && strings.Contains(filePath, "cluster-ae") {
				zone = "europe-west1-b"
			} else if zone == "europe-west1-d" && strings.Contains(filePath, "cluster-sa") {
				zone = "europe-west1-c"
			} else if zone == "us-east4-c" && strings.Contains(filePath, "cluster-cl") {
				zone = "us-east4-b"
			}
			isZoneFound = true
		} else if !isClusterNameFound && rxClusterName.MatchString(scanner.Text()) {
			if code == "qat" && version == 1 {
				//skip version 1 cluster name
				version = 2
			} else {
				clusterName = extractValueWithinDoubleQuotes(scanner.Text())
				isClusterNameFound = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if env == "stg" {
		if code == "kw" {
			code = "stg"
		} else if code == "lt" {
			code = "test"
		} else {
			log.Fatalf("%s:%s", filePath, code)
		}
	}
	return code, region, zone, clusterName
}

var rxVar = regexp.MustCompile("\\$\\{(.*?)\\}")

func formClusterName(clusterNameTemplate, code, region, zone string) string {
	tempCode := code
	if tempCode == "qat" {
		tempCode = "qa"
	}
	regionVar := fmt.Sprintf("${local.%s_region}", tempCode)
	countryCodeVar := fmt.Sprintf("${local.%s_country_code}", tempCode)

	if strings.Contains(clusterNameTemplate, regionVar) {
		clusterNameTemplate = strings.Replace(clusterNameTemplate, regionVar, region, -1)
	}
	if strings.Contains(clusterNameTemplate, countryCodeVar) {
		clusterNameTemplate = strings.Replace(clusterNameTemplate, countryCodeVar, code, -1)
	}
	return clusterNameTemplate
}

func extractCode(fileName string) string {
	fileName = strings.Replace(fileName, "cluster-", "", -1)
	code := strings.Replace(fileName, ".tf", "", -1)

	if code == "kw" {
		code = "kwt"
	} else if code == "qa" {
		code = "qat"
	} else if code == "ar" {
		code = "arg" //Not needed though since file name has arg
	}

	return code
}

func extractCountryCode(code string) string {
	if code == "stg" || code == "test" {
		return code
	}
	return code[:2]
}

func extractCountryInfo(env, path string) map[string]countryInfo {
	countryInfos := make(map[string]countryInfo)

	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if !item.IsDir() && strings.Contains(item.Name(), "cluster") {
			filePath := path + "/" + item.Name()
			code, region, zone, clusterNameTemplate := extractInfraInfo(env, filePath)
			c := countryInfo{
				code:        code,
				countryCode: extractCountryCode(code),
				filePath:    filePath,
				region:      region,
				zone:        zone,
				clusterName: formClusterName(clusterNameTemplate, code, region, zone),
			}
			countryInfos[c.countryCode] = c
		}
	}

	return countryInfos
}

func formGkeClusterName(env, zone, clusterName string) string {
	return fmt.Sprintf("gke_dh-darkstores-%s_%s_%s", env, zone, clusterName)
}

func formDB(env, region, code string) string {
	// Example:
	//        "dh-darkstores-stg:europe-west1:cloudsql-stg-europe-v2=tcp:54325" # stg
	//        "dh-darkstores-live:asia-southeast1:cloudsql-live-tw${REPLICA}=tcp:54322" # tw
	//        "dh-darkstores-live:europe-west1:cloudsql-live-kwt${REPLICA}=tcp:54323" # kw
	prefix := fmt.Sprintf("dh-darkstores-%s:%s:cloudsql-%s-", env, region, env)
	postfix := code
	if code == "pk" {
		postfix += "-v1"
	}
	if env == "stg" {
		postfix = "europe-v2"
	}
	return prefix + postfix
}

func print(env, path string) {
	infos := extractCountryInfo(env, path)
	for _, cc := range countryCodes {
		info, ok := infos[strings.ToLower(cc)]
		if ok {
			fmt.Printf("%s|%s|%s|%s|%s|%s|%s|%d|%s\n",
				info.code,
				info.countryCode,
				countryName[strings.ToUpper(info.countryCode)],
				info.region,
				info.zone,
				info.clusterName,
				formGkeClusterName(env, info.zone, info.clusterName),
				portNumbers[strings.ToUpper(info.countryCode)],
				formDB(env, info.region, info.code),
			)
		}
	}
}

func run() {
	fmt.Printf("1-code|2-country-code|3-country|4-region|5-zone|6-cluster-name|7-gke-cluster-name|8-port|9-db\n")
	print("live", infraPath()+"/terraform/live/infra")
	print("stg", infraPath()+"/terraform/stg/infra")
}

func main() {
	run()
}
