package locales

var Sv = map[string]interface{}{
	"name": map[string]interface{}{
		"first_name_men": []string{
			"Erik", "Lars", "Karl", "Anders", "Per", "Johan", "Nils", "Lennart", "Emil", "Hans",
		},
		"last_name": []string{
			"Johansson", "Andersson", "Karlsson", "Nilsson", "Eriksson", "Larsson", "Olsson", "Persson", "Svensson", "Gustafsson",
		},
		"prefix": []string{
			"Dr.", "Prof.", "PhD.",
		},
		"title": map[string]interface{}{
			"job": []string{
				"Supervisor", "Associate", "Executive", "Liason", "Officer", "Manager", "Engineer", "Specialist", "Director", "Coordinator", "Administrator", "Architect", "Analyst", "Designer", "Planner", "Orchestrator", "Technician", "Developer", "Producer", "Consultant", "Assistant", "Facilitator", "Agent", "Representative", "Strategist",
			},
			"descriptor": []string{
				"Lead", "Senior", "Direct", "Corporate", "Dynamic", "Future", "Product", "National", "Regional", "District", "Central", "Global", "Customer", "Investor", "Dynamic", "International", "Legacy", "Forward", "Internal", "Human", "Chief", "Principal",
			},
			"level": []string{
				"Solutions", "Program", "Brand", "Security", "Research", "Marketing", "Directives", "Implementation", "Integration", "Functionality", "Response", "Paradigm", "Tactics", "Identity", "Markets", "Group", "Division", "Applications", "Optimization", "Operations", "Infrastructure", "Intranet", "Communications", "Web", "Branding", "Quality", "Assurance", "Mobility", "Accounts", "Data", "Creative", "Configuration", "Accountability", "Interactions", "Factors", "Usability", "Metrics",
			},
		},
		"name": []string{
			"#{name.first_name_women} #{name.last_name}", "#{name.first_name_men} #{name.last_name}", "#{name.first_name_women} #{name.last_name}", "#{name.first_name_men} #{name.last_name}", "#{name.first_name_women} #{name.last_name}", "#{name.first_name_men} #{name.last_name}", "#{name.prefix} #{name.first_name_men} #{name.last_name}", "#{name.prefix} #{name.first_name_women} #{name.last_name}",
		},
		"first_name_women": []string{
			"Maria", "Anna", "Margareta", "Elisabeth", "Eva", "Birgitta", "Kristina", "Karin", "Elisabet", "Marie",
		},
	},
	"phone_number": map[string]interface{}{
		"formats": []string{
			"####-#####", "####-######",
		},
	},
	"cell_phone": map[string]interface{}{
		"common_cell_prefix": []string{
			"070", "076", "073",
		},
		"formats": []string{
			"#{common_cell_prefix}-###-####",
		},
	},
	"commerce": map[string]interface{}{
		"color": []string{
			"vit", "silver", "gr??", "svart", "r??d", "gr??n", "bl??", "gul", "lila", "indigo", "guld", "brun", "rosa", "purpur", "korall",
		},
		"department": []string{
			"B??cker", "Filmer", "Musik", "Spel", "Elektronik", "Datorer", "Hem", "Tr??dg??rd", "Verktyg", "Livsmedel", "H??lsa", "Sk??nhet", "Leksaker", "Kl??dsel", "Skor", "Smycken", "Sport",
		},
		"product_name": map[string]interface{}{
			"adjective": []string{
				"Liten", "Ergonomisk", "Robust", "Intelligent", "S??t", "Otrolig", "Fatastisk", "Praktisk", "Slimmad", "Grym",
			},
			"material": []string{
				"St??l", "Metall", "Tr??", "Betong", "Plast", "Bomul", "Grnit", "Gummi", "Latex",
			},
			"product": []string{
				"Stol", "Bil", "Dator", "Handskar", "Pants", "Shirt", "Table", "Shoes", "Hat",
			},
		},
	},
	"team": map[string]interface{}{
		"suffix": []string{
			"IF", "FF", "BK", "HK", "AIF", "SK", "FC", "SK", "BoIS", "FK", "BIS", "FIF", "IK",
		},
		"name": []string{
			"#{address.city} #{name.suffix}",
		},
	},
	"address": map[string]interface{}{
		"street_prefix": []string{
			"V??stra", "??stra", "Norra", "S??dra", "??vre", "Undre",
		},
		"city": []string{
			"#{address.city_prefix}#{address.city_suffix}",
		},
		"secondary_address": []string{
			"Lgh. ###", "Hus ###",
		},
		"street_name": []string{
			"#{street_root}#{address.street_suffix}", "#{street_prefix} #{street_root}#{address.street_suffix}", "#{name.first_name}#{common_street_suffix}", "#{name.last_name}#{common_street_suffix}",
		},
		"street_suffix": []string{
			"v??gen", "gatan", "gr??nden", "g??rdet", "all??n",
		},
		"postcode": []string{
			"#####",
		},
		"default_country": []string{
			"Sverige",
		},
		"city_prefix": []string{
			"S??der", "Norr", "V??st", "??ster", "Aling", "Ar", "Av", "Bo", "Br", "B??", "Ek", "En", "Esk", "Fal", "G??v", "G??te", "Ha", "Helsing", "Karl", "Krist", "Kram", "Kung", "K??", "Lyck", "Ny",
		},
		"city_suffix": []string{
			"stad", "land", "s??s", "??s", "holm", "tuna", "sta", "berg", "l??v", "borg", "mora", "hamn", "fors", "k??ping", "by", "hult", "torp", "fred", "vik",
		},
		"country": []string{
			"Ryssland", "Kanada", "Kina", "USA", "Brasilien", "Australien", "Indien", "Argentina", "Kazakstan", "Algeriet", "DR Kongo", "Danmark", "F??r??arna", "Gr??nland", "Saudiarabien", "Mexiko", "Indonesien", "Sudan", "Libyen", "Iran", "Mongoliet", "Peru", "Tchad", "Niger", "Angola", "Mali", "Sydafrika", "Colombia", "Etiopien", "Bolivia", "Mauretanien", "Egypten", "Tanzania", "Nigeria", "Venezuela", "Namibia", "Pakistan", "Mo??ambique", "Turkiet", "Chile", "Zambia", "Marocko", "V??stsahara", "Burma", "Afghanistan", "Somalia", "Centralafrikanska republiken", "Sydsudan", "Ukraina", "Botswana", "Madagaskar", "Kenya", "Frankrike", "Franska Guyana", "Jemen", "Thailand", "Spanien", "Turkmenistan", "Kamerun", "Papua Nya Guinea", "Sverige", "Uzbekistan", "Irak", "Paraguay", "Zimbabwe", "Japan", "Tyskland", "Kongo", "Finland", "Malaysia", "Vietnam", "Norge", "Svalbard", "Jan Mayen", "Elfenbenskusten", "Polen", "Italien", "Filippinerna", "Ecuador", "Burkina Faso", "Nya Zeeland", "Gabon", "Guinea", "Storbritannien", "Ghana", "Rum??nien", "Laos", "Uganda", "Guyana", "Oman", "Vitryssland", "Kirgizistan", "Senegal", "Syrien", "Kambodja", "Uruguay", "Tunisien", "Surinam", "Nepal", "Bangladesh", "Tadzjikistan", "Grekland", "Nicaragua", "Eritrea", "Nordkorea", "Malawi", "Benin", "Honduras", "Liberia", "Bulgarien", "Kuba", "Guatemala", "Island", "Sydkorea", "Ungern", "Portugal", "Jordanien", "Serbien", "Azerbajdzjan", "??sterrike", "F??renade Arabemiraten", "Tjeckien", "Panama", "Sierra Leone", "Irland", "Georgien", "Sri Lanka", "Litauen", "Lettland", "Togo", "Kroatien", "Bosnien och Hercegovina", "Costa Rica", "Slovakien", "Dominikanska republiken", "Bhutan", "Estland", "Danmark", "F??r??arna", "Gr??nland", "Nederl??nderna", "Schweiz", "Guinea-Bissau", "Taiwan", "Moldavien", "Belgien", "Lesotho", "Armenien", "Albanien", "Salomon??arna", "Ekvatorialguinea", "Burundi", "Haiti", "Rwanda", "Makedonien", "Djibouti", "Belize", "Israel", "El Salvador", "Slovenien", "Fiji", "Kuwait", "Swaziland", "Timor-Leste", "Montenegro", "Bahamas", "Vanuatu", "Qatar", "Gambia", "Jamaica", "Kosovo", "Libanon", "Cypern", "Brunei", "Trinidad och Tobago", "Kap Verde", "Samoa", "Luxemburg", "Komorerna", "Mauritius", "S??o Tom?? och Pr??ncipe", "Kiribati", "Dominica", "Tonga", "Mikronesiens federerade stater", "Singapore", "Bahrain", "Saint Lucia", "Andorra", "Palau", "Seychellerna", "Antigua och Barbuda", "Barbados", "Saint Vincent och Grenadinerna", "Grenada", "Malta", "Maldiverna", "Saint Kitts och Nevis", "Marshall??arna", "Liechtenstein", "San Marino", "Tuvalu", "Nauru", "Monaco", "Vatikanstaten",
		},
		"common_street_suffix": []string{
			"s V??g", "s Gata",
		},
		"street_root": []string{
			"Bj??rk", "J??rnv??gs", "Ring", "Skol", "Skogs", "Ny", "Gran", "Idrotts", "Stor", "Kyrk", "Industri", "Park", "Strand", "Skol", "Tr??dg??rd", "??ngs", "Kyrko", "Villa", "Ek", "Kvarn", "Stations", "Back", "Furu", "Gen", "Fabriks", "??ker", "B??ck", "Asp",
		},
		"state": []string{
			"Blekinge", "Dalarna", "Gotland", "G??vleborg", "G??teborg", "Halland", "J??mtland", "J??nk??ping", "Kalmar", "Kronoberg", "Norrbotten", "Skaraborg", "Sk??ne", "Stockholm", "S??dermanland", "Uppsala", "V??rmland", "V??sterbotten", "V??sternorrland", "V??stmanland", "??lvsborg", "??rebro", "??sterg??tland",
		},
		"building_number": []string{
			"###", "##", "#",
		},
		"street_address": []string{
			"#{address.street_name} #{address.building_number}",
		},
	},
	"company": map[string]interface{}{
		"suffix": []string{
			"Gruppen", "AB", "HB", "Group", "Investment", "Kommanditbolag", "Aktiebolag",
		},
		"name": []string{
			"#{name.last_name} #{name.suffix}", "#{name.last_name}-#{name.last_name}", "#{name.last_name}, #{name.last_name} #{name.suffix}",
		},
	},
	"internet": map[string]interface{}{
		"domain_suffix": []string{
			"se", "nu", "info", "com", "org"}}}
