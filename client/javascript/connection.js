cb.connection = function() {
"use strict";
	var networks = [
		{
			name: "freenode",
			address: "irc.freenode.net",
			nickname: "Mal",
			channels: [
				{
					name: "#css",
					topic: "Cascading Style Sheets :: Don't paste! :: Need help? Validate (`v yoursite) and give a `live URL (see `paste and `testcase) :: Rules: http://hashcss.com/rules/ :: FAQ: http://hashcss.com/faq :: Harassment or disruption? Tell #css-ops :: More general discussions: ##frontend",
					icon: "http://i.imgur.com/ORe6fC1.png",
					users: [
						"Keavon",
						"Joshua-Anderson"
					],
					messages: {
						
					}
				}, {
					name: "##javascript",
					topic: "Can't talk? Get registered on freenode (HOWTO: http://freenode.net/faq.shtml#nicksetup ). | ECMAScript, JavaScript. JS *not* Java. | Say \"!help\" (or ask and wait). | Say \"!mdn abc\" for docs on \"abc\". | Don't paste code in the channel.",
					icon: "http://i.imgur.com/dp4NsvM.png",
					users: [
						"Mal",
						"Wash",
						"Zoey",
						"Jayne",
						"Inara",
						"Book",
						"Kayle",
						"Simon",
						"River"
					],
					messages: {
						
					}
				}, {
					name: "#html",
					topic: "HyperText Markup Language | http://www.w3.org/TR/html5/ | #css for styling | No WYSIWYGs | Tables aren't for layout: http://xrl.us/kd35 , http://goo.gl/8yFKr | Separate Content and Presentation: http://goo.gl/2DtpJO | Validate first: http://validator.w3.org/ | W3C reference: say e.g. `html5 div | Paste links, not code: `paste",
					icon: "http://i.imgur.com/V2CKhc3.png",
					users: [
						"Mal",
						"Wash",
						"Zoey",
						"Jayne",
						"Inara",
						"Book",
						"Kayle",
						"Simon",
						"River"
					],
					messages: {
						
					}
				}, {
					name: "#blender",
					topic: "Blender Support Channel | www.blender.org | Current release 2.75a | Please avoid swearing | Share files with www.pasteall.org | More Q&A at blender.stackexchange.com | Check out the Blender Cloud and help support future projects: https://cloud.blender.org",
					icon: "http://i.imgur.com/0XMMtNp.png",
					users: [
						"Mal",
						"Wash",
						"Zoey",
						"Jayne",
						"Inara",
						"Book",
						"Kayle",
						"Simon",
						"River"
					],
					messages: {
						
					}
				}
			]
		}
	];
	
	return {
		networks: networks
	};
}();
