// Run when DOM loads
document.addEventListener('DOMContentLoaded', function() {
	// Select tab bar element
	var tabBar = document.getElementsByClassName("tabs")[0];

	// Add a tab for each channel
	for (var server in servers) {
		for (var channel in servers[server].channels) {
			channel = servers[server].channels[channel];

			// Create tab element
			var tab = document.createElement("a");

			// Include inner HTML for tab with channel name and icon
			tab.innerHTML = "<span><img src=\"" + channel.icon + "\">" + channel.name + "</span><svg viewBox=\"0 0 1000 1000\"><circle cx=\"500\" cy=\"500\" r=\"500\"></circle><line x1=\"250\" y1=\"750\" x2=\"750\" y2=\"250\"></line><line x1=\"250\" y1=\"250\" x2=\"750\" y2=\"750\"></line></svg>";

			// If tab is the currently selected one, include `active` class
			if (channel.name === currentChannel) {
				tab.className = "active";
			}

			// Append tab to tab bar
			tabBar.appendChild(tab);
		}
	}

	// Append new tab button
	tabBar.innerHTML += "<a class=\"add\"><svg viewBox=\"0 0 1000 1000\"><path d=\"M791.7,541.7h-250v250h-83.3v-250h-250v-83.3h250v-250h83.3v250h250V541.7z\" /></svg></a>";

	// Add click listener to tab buttons
	tabBar.addEventListener("click", function(event) {
		// Check if click is on the new tab button
		if (event.target.className === "add" || event.target.parentElement.className === "add" || event.target.parentElement.parentElement.className === "add") {
			addTab();
			return;
		}

		// Check if click is on a close tab button
		if (event.target.nodeName === "svg") {
			closeTab(event.target.parentElement.childNodes[0].childNodes[1].nodeValue);
			return;
		} else if (event.target.parentElement.nodeName === "svg") {
			closeTab(event.target.parentElement.parentElement.childNodes[0].childNodes[1].nodeValue);
			return;
		}

		// Check if click is on a tab
		if (event.target.nodeName.toLowerCase() === "a") {
			changeTab(event.target.childNodes[0].childNodes[1].nodeValue);
			return;
		} else if (event.target.parentElement.nodeName.toLowerCase() === "a") {
			changeTab(event.target.childNodes[1].nodeValue);
			return;
		} else if (event.target.parentElement.parentElement.nodeName.toLowerCase() === "a") {
			changeTab(event.target.parentElement.childNodes[1].nodeValue);
			return;
		}
	});

	// Initially load the default tab
	changeTab(currentChannel);
});

var changeTab = function(tabName) {
	// Find channel from array of channels in all servers
	var channel;
	var server;
	servers.some(function(s) {
		return s.channels.some(function(c) {
			if (c.name === tabName) {
				channel = c;
				server = s;
				return true;
			}
		});
	});

	// Remove existing active tab
	if (document.querySelector(".tabs .active")) {
		document.querySelector(".tabs .active").removeAttribute("class");
	}

	// Select each tab element
	var tabs = document.getElementsByClassName("tabs")[0].childNodes;

	// Make the correct tab active
	for (var tab in tabs) {
		// Throw out the ones that aren't just tabs
		if (tabs[tab].className === "") {
			// Check if this tab has the text with the desired tab name
			if (tabs[tab].childNodes[0].childNodes[1].nodeValue === channel.name) {
				// Set this tab as active
				tabs[tab].className = "active";
			}
		}
	}

	// Change channel icon, name, topic, user count, and server name and nickname
	document.querySelector(".channel .info img").src = channel.icon;
	document.querySelector(".channel .info .title h1").innerHTML = channel.name;
	document.querySelector(".channel .info .title h6").innerHTML = server.name;
	document.querySelector(".channel p").innerHTML = channel.topic;
	document.querySelector(".channel .info .users span").innerHTML = channel.users.length;
	document.querySelector(".input .name a").innerHTML = server.nickname;

	// Add user list to users sidebar
	var users = "";
	for (var name in channel.users) {
		users += "<li><a href=\"#\">" + channel.users[name] + "</a></li>";
	}
	document.querySelector(".user-list ul").innerHTML = users;
};

var closeTab = function(tabName) {
	// Select each tab
	var tabs = document.querySelectorAll(".tabs a");

	// Find the tab with the matching name and remove it
	for (var tab in tabs) {
		// Make sure the element is an actual tab
		if (tabs[tab].className === "" || tabs[tab].className === "active") {
			// Check if the tab text matches the tab to be closed
			if (tabs[tab].childNodes[0].childNodes[1].nodeValue === tabName) {
				// Remove the tab from the parent
				document.getElementsByClassName("tabs")[0].removeChild(tabs[tab]);
			}
		}
	}
};

var addTab = function() {
	alert("Adding tab");
};
