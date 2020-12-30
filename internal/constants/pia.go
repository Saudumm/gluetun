package constants

import (
	"net"

	"github.com/qdm12/gluetun/internal/models"
)

//nolint:lll
const (
	PIAEncryptionPresetNormal = "normal"
	PIAEncryptionPresetStrong = "strong"
	PiaX509CRLNormal          = "MIICWDCCAUAwDQYJKoZIhvcNAQENBQAwgegxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTETMBEGA1UEBxMKTG9zQW5nZWxlczEgMB4GA1UEChMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAsTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQDExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEKRMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxLzAtBgkqhkiG9w0BCQEWIHNlY3VyZUBwcml2YXRlaW50ZXJuZXRhY2Nlc3MuY29tFw0xNjA3MDgxOTAwNDZaFw0zNjA3MDMxOTAwNDZaMCYwEQIBARcMMTYwNzA4MTkwMDQ2MBECAQYXDDE2MDcwODE5MDA0NjANBgkqhkiG9w0BAQ0FAAOCAQEAQZo9X97ci8EcPYu/uK2HB152OZbeZCINmYyluLDOdcSvg6B5jI+ffKN3laDvczsG6CxmY3jNyc79XVpEYUnq4rT3FfveW1+Ralf+Vf38HdpwB8EWB4hZlQ205+21CALLvZvR8HcPxC9KEnev1mU46wkTiov0EKc+EdRxkj5yMgv0V2Reze7AP+NQ9ykvDScH4eYCsmufNpIjBLhpLE2cuZZXBLcPhuRzVoU3l7A9lvzG9mjA5YijHJGHNjlWFqyrn1CfYS6koa4TGEPngBoAziWRbDGdhEgJABHrpoaFYaL61zqyMR6jC0K2ps9qyZAN74LEBedEfK7tBOzWMwr58A=="
	PiaX509CRLStrong          = "MIIDWDCCAUAwDQYJKoZIhvcNAQENBQAwgegxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTETMBEGA1UEBxMKTG9zQW5nZWxlczEgMB4GA1UEChMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAsTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQDExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEKRMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxLzAtBgkqhkiG9w0BCQEWIHNlY3VyZUBwcml2YXRlaW50ZXJuZXRhY2Nlc3MuY29tFw0xNjA3MDgxOTAwNDZaFw0zNjA3MDMxOTAwNDZaMCYwEQIBARcMMTYwNzA4MTkwMDQ2MBECAQYXDDE2MDcwODE5MDA0NjANBgkqhkiG9w0BAQ0FAAOCAgEAppFfEpGsasjB1QgJcosGpzbf2kfRhM84o2TlqY1ua+Gi5TMdKydA3LJcNTjlI9a0TYAJfeRX5IkpoglSUuHuJgXhP3nEvX10mjXDpcu/YvM8TdE5JV2+EGqZ80kFtBeOq94WcpiVKFTR4fO+VkOK9zwspFfb1cNs9rHvgJ1QMkRUF8PpLN6AkntHY0+6DnigtSaKqldqjKTDTv2OeH3nPoh80SGrt0oCOmYKfWTJGpggMGKvIdvU3vH9+EuILZKKIskt+1dwdfA5Bkz1GLmiQG7+9ZZBQUjBG9Dos4hfX/rwJ3eU8oUIm4WoTz9rb71SOEuUUjP5NPy9HNx2vx+cVvLsTF4ZDZaUztW9o9JmIURDtbeyqxuHN3prlPWB6aj73IIm2dsDQvs3XXwRIxs8NwLbJ6CyEuvEOVCskdM8rdADWx1J0lRNlOJ0Z8ieLLEmYAA834VN1SboB6wJIAPxQU3rcBhXqO9y8aa2oRMg8NxZ5gr+PnKVMqag1x0IxbIgLxtkXQvxXxQHEMSODzvcOfK/nBRBsqTj30P+R87sU8titOoxNeRnBDRNhdEy/QGAqGh62ShPpQUCJdnKRiRTjnil9hMQHevoSuFKeEMO30FQL7BZyo37GFU+q1WPCplVZgCP9hC8Rn5K2+f6KLFo5bhtowSmu+GY1yZtg+RTtsA="
	PIACertificateNormal      = "MIIFqzCCBJOgAwIBAgIJAKZ7D5Yv87qDMA0GCSqGSIb3DQEBDQUAMIHoMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExEzARBgNVBAcTCkxvc0FuZ2VsZXMxIDAeBgNVBAoTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQLExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEAxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBCkTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMS8wLQYJKoZIhvcNAQkBFiBzZWN1cmVAcHJpdmF0ZWludGVybmV0YWNjZXNzLmNvbTAeFw0xNDA0MTcxNzM1MThaFw0zNDA0MTIxNzM1MThaMIHoMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExEzARBgNVBAcTCkxvc0FuZ2VsZXMxIDAeBgNVBAoTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQLExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEAxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBCkTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMS8wLQYJKoZIhvcNAQkBFiBzZWN1cmVAcHJpdmF0ZWludGVybmV0YWNjZXNzLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPXDL1L9tX6DGf36liA7UBTy5I869z0UVo3lImfOs/GSiFKPtInlesP65577nd7UNzzXlH/P/CnFPdBWlLp5ze3HRBCc/Avgr5CdMRkEsySL5GHBZsx6w2cayQ2EcRhVTwWpcdldeNO+pPr9rIgPrtXqT4SWViTQRBeGM8CDxAyTopTsobjSiYZCF9Ta1gunl0G/8Vfp+SXfYCC+ZzWvP+L1pFhPRqzQQ8k+wMZIovObK1s+nlwPaLyayzw9a8sUnvWB/5rGPdIYnQWPgoNlLN9HpSmsAcw2z8DXI9pIxbr74cb3/HSfuYGOLkRqrOk6h4RCOfuWoTrZup1uEOn+fw8CAwEAAaOCAVQwggFQMB0GA1UdDgQWBBQv63nQ/pJAt5tLy8VJcbHe22ZOsjCCAR8GA1UdIwSCARYwggESgBQv63nQ/pJAt5tLy8VJcbHe22ZOsqGB7qSB6zCB6DELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMRMwEQYDVQQHEwpMb3NBbmdlbGVzMSAwHgYDVQQKExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UECxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAMTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQpExdQcml2YXRlIEludGVybmV0IEFjY2VzczEvMC0GCSqGSIb3DQEJARYgc2VjdXJlQHByaXZhdGVpbnRlcm5ldGFjY2Vzcy5jb22CCQCmew+WL/O6gzAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBDQUAA4IBAQAna5PgrtxfwTumD4+3/SYvwoD66cB8IcK//h1mCzAduU8KgUXocLx7QgJWo9lnZ8xUryXvWab2usg4fqk7FPi00bED4f4qVQFVfGfPZIH9QQ7/48bPM9RyfzImZWUCenK37pdw4Bvgoys2rHLHbGen7f28knT2j/cbMxd78tQc20TIObGjo8+ISTRclSTRBtyCGohseKYpTS9himFERpUgNtefvYHbn70mIOzfOJFTVqfrptf9jXa9N8Mpy3ayfodz1wiqdteqFXkTYoSDctgKMiZ6GdocK9nMroQipIQtpnwd4yBDWIyC6Bvlkrq5TQUtYDQ8z9v+DMO6iwyIDRiU"
	PIACertificateStrong      = "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"
)

func PIAGeoChoices() (choices []string) {
	servers := PIAServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return choices
}

//nolint:lll
func PIAServers() []models.PIAServer {
	return []models.PIAServer{
		{Region: "AU Melbourne", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "melbourne404", IPs: []net.IP{{103, 2, 198, 93}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "melbourne404", IPs: []net.IP{{103, 2, 198, 85}}}},
		{Region: "AU Perth", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "perth403", IPs: []net.IP{{43, 250, 205, 156}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "perth403", IPs: []net.IP{{43, 250, 205, 153}}}},
		{Region: "AU Sydney", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "sydney408", IPs: []net.IP{{117, 120, 10, 56}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "sydney408", IPs: []net.IP{{117, 120, 10, 48}}}},
		{Region: "Albania", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "tirana402", IPs: []net.IP{{31, 171, 154, 115}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "tirana402", IPs: []net.IP{{31, 171, 154, 118}}}},
		{Region: "Algeria", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "algiers404", IPs: []net.IP{{176, 125, 228, 19}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "algiers404", IPs: []net.IP{{176, 125, 228, 26}}}},
		{Region: "Andorra", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "andorra405", IPs: []net.IP{{188, 241, 82, 40}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "andorra405", IPs: []net.IP{{188, 241, 82, 37}}}},
		{Region: "Argentina", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "buenosaires401", IPs: []net.IP{{190, 106, 134, 95}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "buenosaires401", IPs: []net.IP{{190, 106, 134, 85}}}},
		{Region: "Armenia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "armenia403", IPs: []net.IP{{185, 253, 160, 3}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "armenia403", IPs: []net.IP{{185, 253, 160, 9}}}},
		{Region: "Austria", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "vienna403", IPs: []net.IP{{156, 146, 60, 108}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "vienna403", IPs: []net.IP{{156, 146, 60, 110}}}},
		{Region: "Bahamas", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "bahamas404", IPs: []net.IP{{95, 181, 238, 9}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "bahamas404", IPs: []net.IP{{95, 181, 238, 9}}}},
		{Region: "Bangladesh", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "bangladesh403", IPs: []net.IP{{84, 252, 93, 7}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "bangladesh403", IPs: []net.IP{{84, 252, 93, 7}}}},
		{Region: "Belgium", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "brussels404", IPs: []net.IP{{37, 120, 143, 156}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "brussels404", IPs: []net.IP{{37, 120, 143, 156}}}},
		{Region: "Bosnia and Herzegovina", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "sarajevo402", IPs: []net.IP{{185, 212, 111, 4}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "sarajevo402", IPs: []net.IP{{185, 212, 111, 71}}}},
		{Region: "Brazil", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "saopaolo401", IPs: []net.IP{{45, 133, 180, 236}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "saopaolo401", IPs: []net.IP{{45, 133, 180, 227}}}},
		{Region: "Bulgaria", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "sofia402", IPs: []net.IP{{217, 138, 221, 89}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "sofia402", IPs: []net.IP{{217, 138, 221, 89}}}},
		{Region: "CA Montreal", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "montreal410", IPs: []net.IP{{199, 36, 223, 210}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "montreal410", IPs: []net.IP{{199, 36, 223, 210}}}},
		{Region: "CA Ontario", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "ontario408", IPs: []net.IP{{172, 98, 92, 87}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "ontario408", IPs: []net.IP{{172, 98, 92, 80}}}},
		{Region: "CA Toronto", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "toronto402", IPs: []net.IP{{66, 115, 142, 58}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "toronto402", IPs: []net.IP{{66, 115, 142, 58}}}},
		{Region: "CA Vancouver", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "vancouver411", IPs: []net.IP{{208, 78, 42, 168}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "vancouver411", IPs: []net.IP{{208, 78, 42, 161}}}},
		{Region: "Cambodia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "cambodia401", IPs: []net.IP{{188, 215, 235, 109}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "cambodia401", IPs: []net.IP{{188, 215, 235, 110}}}},
		{Region: "China", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "china404", IPs: []net.IP{{188, 241, 80, 9}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "china404", IPs: []net.IP{{188, 241, 80, 4}}}},
		{Region: "Cyprus", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "cyprus403", IPs: []net.IP{{185, 253, 162, 8}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "cyprus403", IPs: []net.IP{{185, 253, 162, 14}}}},
		{Region: "Czech Republic", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "prague405", IPs: []net.IP{{143, 244, 59, 168}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "prague405", IPs: []net.IP{{143, 244, 59, 154}}}},
		{Region: "DE Berlin", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "berlin425", IPs: []net.IP{{154, 13, 1, 148}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "berlin425", IPs: []net.IP{{154, 13, 1, 146}}}},
		{Region: "DE Frankfurt", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "frankfurt440", IPs: []net.IP{{185, 216, 33, 165}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "frankfurt440", IPs: []net.IP{{185, 216, 33, 166}}}},
		{Region: "Denmark", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "copenhagen404", IPs: []net.IP{{188, 126, 94, 190}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "copenhagen404", IPs: []net.IP{{188, 126, 94, 165}}}},
		{Region: "Egypt", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "cairo402", IPs: []net.IP{{188, 214, 122, 126}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "cairo402", IPs: []net.IP{{188, 214, 122, 123}}}},
		{Region: "Estonia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "talinn402", IPs: []net.IP{{95, 153, 31, 68}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "talinn402", IPs: []net.IP{{95, 153, 31, 78}}}},
		{Region: "Finland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "helsinki402", IPs: []net.IP{{188, 126, 89, 45}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "helsinki402", IPs: []net.IP{{188, 126, 89, 35}}}},
		{Region: "France", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "paris406", IPs: []net.IP{{143, 244, 57, 169}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "paris406", IPs: []net.IP{{143, 244, 57, 169}}}},
		{Region: "Georgia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "georgia403", IPs: []net.IP{{95, 181, 236, 8}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "georgia403", IPs: []net.IP{{95, 181, 236, 10}}}},
		{Region: "Greece", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "athens401", IPs: []net.IP{{154, 57, 3, 87}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "athens401", IPs: []net.IP{{154, 57, 3, 85}}}},
		{Region: "Greenland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "greenland404", IPs: []net.IP{{91, 90, 120, 149}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "greenland404", IPs: []net.IP{{91, 90, 120, 147}}}},
		{Region: "Hong Kong", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "china403", IPs: []net.IP{{86, 107, 104, 213}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "china403", IPs: []net.IP{{86, 107, 104, 219}}}},
		{Region: "Hungary", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "budapest402", IPs: []net.IP{{86, 106, 74, 115}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "budapest402", IPs: []net.IP{{86, 106, 74, 117}}}},
		{Region: "Iceland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "reykjavik402", IPs: []net.IP{{45, 133, 193, 83}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "reykjavik402", IPs: []net.IP{{45, 133, 193, 88}}}},
		{Region: "India", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "mumbai402", IPs: []net.IP{{45, 120, 139, 138}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "mumbai402", IPs: []net.IP{{45, 120, 139, 127}}}},
		{Region: "Ireland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "dublin411", IPs: []net.IP{{188, 241, 178, 23}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "dublin411", IPs: []net.IP{{188, 241, 178, 30}}}},
		{Region: "Isle of Man", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "douglas403", IPs: []net.IP{{91, 90, 124, 7}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "douglas403", IPs: []net.IP{{91, 90, 124, 18}}}},
		{Region: "Israel", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "jerusalem407", IPs: []net.IP{{185, 77, 248, 91}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "jerusalem407", IPs: []net.IP{{185, 77, 248, 91}}}},
		{Region: "Italy", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "milano403", IPs: []net.IP{{156, 146, 41, 74}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "milano403", IPs: []net.IP{{156, 146, 41, 84}}}},
		{Region: "Japan", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "tokyo401", IPs: []net.IP{{156, 146, 34, 130}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "tokyo401", IPs: []net.IP{{156, 146, 34, 130}}}},
		{Region: "Kazakhstan", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "kazakhstan403", IPs: []net.IP{{62, 133, 47, 13}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "kazakhstan403", IPs: []net.IP{{62, 133, 47, 5}}}},
		{Region: "Latvia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "riga401", IPs: []net.IP{{109, 248, 149, 5}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "riga401", IPs: []net.IP{{109, 248, 149, 8}}}},
		{Region: "Liechtenstein", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "liechtenstein403", IPs: []net.IP{{91, 90, 122, 7}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "liechtenstein403", IPs: []net.IP{{91, 90, 122, 18}}}},
		{Region: "Lithuania", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "vilnius403", IPs: []net.IP{{85, 206, 165, 118}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "vilnius403", IPs: []net.IP{{85, 206, 165, 116}}}},
		{Region: "Luxembourg", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "luxembourg407", IPs: []net.IP{{5, 253, 204, 150}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "luxembourg407", IPs: []net.IP{{5, 253, 204, 155}}}},
		{Region: "Macao", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "macau403", IPs: []net.IP{{84, 252, 92, 6}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "macau403", IPs: []net.IP{{84, 252, 92, 15}}}},
		{Region: "Macedonia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "macedonia401", IPs: []net.IP{{185, 225, 28, 120}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "macedonia401", IPs: []net.IP{{185, 225, 28, 120}}}},
		{Region: "Malta", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "malta403", IPs: []net.IP{{176, 125, 230, 13}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "malta403", IPs: []net.IP{{176, 125, 230, 7}}}},
		{Region: "Mexico", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "mexico403", IPs: []net.IP{{77, 81, 142, 14}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "mexico403", IPs: []net.IP{{77, 81, 142, 13}}}},
		{Region: "Moldova", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "chisinau401", IPs: []net.IP{{178, 175, 129, 44}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "chisinau401", IPs: []net.IP{{178, 175, 129, 46}}}},
		{Region: "Monaco", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "monaco403", IPs: []net.IP{{95, 181, 233, 8}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "monaco403", IPs: []net.IP{{95, 181, 233, 5}}}},
		{Region: "Mongolia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "mongolia403", IPs: []net.IP{{185, 253, 163, 15}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "mongolia403", IPs: []net.IP{{185, 253, 163, 5}}}},
		{Region: "Montenegro", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "montenegro403", IPs: []net.IP{{176, 125, 229, 14}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "montenegro403", IPs: []net.IP{{176, 125, 229, 4}}}},
		{Region: "Morocco", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "morocco403", IPs: []net.IP{{95, 181, 232, 4}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "morocco403", IPs: []net.IP{{95, 181, 232, 8}}}},
		{Region: "Netherlands", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "amsterdam412", IPs: []net.IP{{143, 244, 41, 196}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "amsterdam412", IPs: []net.IP{{143, 244, 41, 196}}}},
		{Region: "New Zealand", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "newzealand403", IPs: []net.IP{{43, 250, 207, 90}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "newzealand403", IPs: []net.IP{{43, 250, 207, 84}}}},
		{Region: "Nigeria", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "nigeria403", IPs: []net.IP{{102, 165, 25, 86}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "nigeria403", IPs: []net.IP{{102, 165, 25, 85}}}},
		{Region: "Norway", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "oslo401", IPs: []net.IP{{46, 246, 122, 37}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "oslo401", IPs: []net.IP{{46, 246, 122, 60}}}},
		{Region: "Panama", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "panama404", IPs: []net.IP{{91, 90, 126, 25}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "panama404", IPs: []net.IP{{91, 90, 126, 28}}}},
		{Region: "Philippines", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "philippines401", IPs: []net.IP{{188, 214, 125, 140}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "philippines401", IPs: []net.IP{{188, 214, 125, 137}}}},
		{Region: "Poland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "warsaw409", IPs: []net.IP{{194, 110, 114, 119}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "warsaw409", IPs: []net.IP{{194, 110, 114, 118}}}},
		{Region: "Portugal", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "lisbon402", IPs: []net.IP{{89, 26, 241, 87}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "lisbon402", IPs: []net.IP{{89, 26, 241, 88}}}},
		{Region: "Qatar", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "qatar403", IPs: []net.IP{{95, 181, 234, 9}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "qatar403", IPs: []net.IP{{95, 181, 234, 8}}}},
		{Region: "Romania", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "romania408", IPs: []net.IP{{143, 244, 54, 117}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "romania408", IPs: []net.IP{{143, 244, 54, 116}}}},
		{Region: "Saudi Arabia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "saudiarabia403", IPs: []net.IP{{95, 181, 235, 8}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "saudiarabia403", IPs: []net.IP{{95, 181, 235, 4}}}},
		{Region: "Serbia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "belgrade401", IPs: []net.IP{{37, 120, 193, 250}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "belgrade401", IPs: []net.IP{{37, 120, 193, 249}}}},
		{Region: "Singapore", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "singapore401", IPs: []net.IP{{156, 146, 57, 213}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "singapore401", IPs: []net.IP{{156, 146, 57, 197}}}},
		{Region: "Slovakia", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "bratislava402", IPs: []net.IP{{37, 120, 221, 213}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "bratislava402", IPs: []net.IP{{37, 120, 221, 218}}}},
		{Region: "South Africa", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "johannesburg401", IPs: []net.IP{{154, 16, 93, 33}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "johannesburg401", IPs: []net.IP{{154, 16, 93, 46}}}},
		{Region: "Spain", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "madrid401", IPs: []net.IP{{212, 102, 49, 68}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "madrid401", IPs: []net.IP{{212, 102, 49, 68}}}},
		{Region: "Sri Lanka", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "srilanka403", IPs: []net.IP{{95, 181, 239, 8}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "srilanka403", IPs: []net.IP{{95, 181, 239, 13}}}},
		{Region: "Sweden", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "stockholm401", IPs: []net.IP{{195, 246, 120, 4}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "stockholm401", IPs: []net.IP{{195, 246, 120, 39}}}},
		{Region: "Switzerland", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "zurich407", IPs: []net.IP{{156, 146, 62, 194}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "zurich407", IPs: []net.IP{{156, 146, 62, 194}}}},
		{Region: "Taiwan", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "taiwan401", IPs: []net.IP{{188, 214, 106, 74}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "taiwan401", IPs: []net.IP{{188, 214, 106, 69}}}},
		{Region: "Turkey", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "istanbul402", IPs: []net.IP{{188, 213, 34, 88}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "istanbul402", IPs: []net.IP{{188, 213, 34, 83}}}},
		{Region: "UK London", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "london405", IPs: []net.IP{{212, 102, 53, 15}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "london405", IPs: []net.IP{{212, 102, 53, 60}}}},
		{Region: "UK London-2", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "london420", IPs: []net.IP{{37, 235, 96, 200}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "london420", IPs: []net.IP{{37, 235, 96, 206}}}},
		{Region: "UK Manchester", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "manchester414", IPs: []net.IP{{194, 37, 96, 194}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "manchester414", IPs: []net.IP{{194, 37, 96, 197}}}},
		{Region: "UK Southampton", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "southampton401", IPs: []net.IP{{143, 244, 37, 244}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "southampton401", IPs: []net.IP{{143, 244, 37, 194}}}},
		{Region: "US Atlanta", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "atlanta417", IPs: []net.IP{{154, 21, 22, 216}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "atlanta417", IPs: []net.IP{{154, 21, 22, 216}}}},
		{Region: "US California", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "losangeles401", IPs: []net.IP{{143, 244, 48, 15}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "losangeles401", IPs: []net.IP{{143, 244, 48, 17}}}},
		{Region: "US Chicago", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "chicago413", IPs: []net.IP{{154, 21, 23, 125}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "chicago413", IPs: []net.IP{{154, 21, 23, 137}}}},
		{Region: "US Denver", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "denver410", IPs: []net.IP{{174, 128, 227, 24}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "denver410", IPs: []net.IP{{174, 128, 227, 24}}}},
		{Region: "US East", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "newjersey405", IPs: []net.IP{{143, 244, 46, 65}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "newjersey405", IPs: []net.IP{{143, 244, 46, 115}}}},
		{Region: "US Florida", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "miami407", IPs: []net.IP{{156, 146, 42, 78}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "miami407", IPs: []net.IP{{156, 146, 42, 113}}}},
		{Region: "US Houston", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "houston413", IPs: []net.IP{{205, 251, 142, 20}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "houston413", IPs: []net.IP{{205, 251, 142, 20}}}},
		{Region: "US Las Vegas", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "lasvegas426", IPs: []net.IP{{196, 53, 64, 175}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "lasvegas426", IPs: []net.IP{{196, 53, 64, 189}}}},
		{Region: "US New York", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "newyork404", IPs: []net.IP{{143, 244, 44, 227}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "newyork404", IPs: []net.IP{{143, 244, 44, 208}}}},
		{Region: "US Seattle", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "seattle422", IPs: []net.IP{{156, 146, 48, 204}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "seattle422", IPs: []net.IP{{156, 146, 48, 239}}}},
		{Region: "US Silicon Valley", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "siliconvalley420", IPs: []net.IP{{66, 115, 165, 93}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "siliconvalley420", IPs: []net.IP{{66, 115, 165, 80}}}},
		{Region: "US Texas", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "dallas415", IPs: []net.IP{{154, 3, 250, 16}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "dallas415", IPs: []net.IP{{154, 3, 250, 24}}}},
		{Region: "US Washington DC", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "washington452", IPs: []net.IP{{91, 149, 244, 110}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "washington452", IPs: []net.IP{{91, 149, 244, 114}}}},
		{Region: "US West", PortForward: false, OpenvpnUDP: models.PIAServerOpenvpn{CN: "phoenix413", IPs: []net.IP{{184, 170, 241, 169}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "phoenix413", IPs: []net.IP{{184, 170, 241, 169}}}},
		{Region: "Ukraine", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "kiev403", IPs: []net.IP{{62, 149, 20, 7}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "kiev403", IPs: []net.IP{{62, 149, 20, 4}}}},
		{Region: "United Arab Emirates", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "dubai403", IPs: []net.IP{{217, 138, 193, 153}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "dubai403", IPs: []net.IP{{217, 138, 193, 150}}}},
		{Region: "Venezuela", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "venezuela403", IPs: []net.IP{{95, 181, 237, 3}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "venezuela403", IPs: []net.IP{{95, 181, 237, 9}}}},
		{Region: "Vietnam", PortForward: true, OpenvpnUDP: models.PIAServerOpenvpn{CN: "vietnam401", IPs: []net.IP{{188, 214, 152, 67}}}, OpenvpnTCP: models.PIAServerOpenvpn{CN: "vietnam401", IPs: []net.IP{{188, 214, 152, 67}}}},
	}
}
