// Stores information relating to tabs
var tabs = {
	list: [],
	active: 0
};

// Run when DOM loads
document.addEventListener('DOMContentLoaded', function() {
	// Add click listener to tab buttons
	document.querySelector(".tabs").addEventListener("click", function(event) {
		// Check if click is on the new tab button
		if (event.target.className === "add" || event.target.parentElement.className === "add" || event.target.parentElement.parentElement.className === "add") {
			newTab();
			return;
		}
		
		// Check if click is on a tab close button
		if (event.target.nodeName === "svg") {
			closeTab(indexInParent(event.target.parentElement));
			return;
		} else if (event.target.parentElement.nodeName === "svg") {
			closeTab(indexInParent(event.target.parentElement.parentElement));
			return;
		}
		
		// Check if click is on a tab
		if (event.target.nodeName.toLowerCase() === "a") {
			changeTab(indexInParent(event.target));
			return;
		} else if (event.target.parentElement.nodeName.toLowerCase() === "a") {
			changeTab(indexInParent(event.target.parentElement));
			return;
		} else if (event.target.parentElement.parentElement.nodeName.toLowerCase() === "a") {
			changeTab(indexInParent(event.target.parentElement.parentElement));
			return;
		}
	});
	
	// Initially open a tab for each channel in each network
	for (var network in networks) {
		for (var channel in networks[network].channels) {
			addTab(networks[network], networks[network].channels[channel]);
		}
	}
	
	// Initially load the default tab
	changeTab(tabs.active);
});

// Switches to a tab given its index
var changeTab = function(tabIndex) {
	// Store the active tab
	tabs.active = tabIndex;
	
	// Remove the active status from the existing active tab
	if (document.querySelector(".tabs .active")) {
		document.querySelector(".tabs .active").removeAttribute("class");
	}

	// Set the new active tab as active
	document.querySelectorAll(".tabs a")[tabIndex].className = "active";
	
	// Store the network and channel for the new tab
	var network = getNetworkFromTabIndex(tabIndex);
	var channel = getChannelFromTabIndex(tabIndex);
	
	// Change channel icon, name, topic, user count, and network name and nickname
	document.querySelector(".channel .info img").src = channel.icon;
	document.querySelector(".channel .info .title h1").innerHTML = channel.name;
	document.querySelector(".channel .info .title h6").innerHTML = network.name;
	document.querySelector(".channel p").innerHTML = channel.topic;
	document.querySelector(".channel .info .users span").innerHTML = channel.users.length;
	document.querySelector(".input .name a").innerHTML = network.nickname;
	
	// Add user list to users sidebar
	var users = "";
	for (var name in channel.users) {
		users += "<li><a href=\"#\">" + channel.users[name] + "</a></li>";
	}
	document.querySelector(".user-list ul").innerHTML = users;
};

// Closes a tab given its index
var closeTab = function(tabIndex) {
	// Prevent the user from closing the last tab
	if (tabs.list.length <= 1) {
		return;
	}
	
	// Remove tab from the list of tabs
	tabs.list.splice(tabIndex, 1);
	
	// Remove tab from the DOM
	document.querySelector(".tabs").removeChild(document.querySelectorAll(".tabs a")[tabIndex]);
	
	// Move to a new tab if the user closes the active tab
	if (tabIndex === tabs.active) {
		if (tabs.active < tabs.list.length) {
			// Move to the one in the new position of the former closed tab
			changeTab(tabs.active);
		} else {
			// Move to the new last tab if the former last tab is closed
			changeTab(tabs.list.length - 1);
		}
	}
};

// Adds a tab for a channel given the network address and channel name
var addTab = function(network, channel) {
	// Create a new tab element
	var tabElement = document.createElement("a");
	
	// Include inner HTML for tab with channel name and icon
	tabElement.innerHTML = "<span><img src=\"" + channel.icon + "\">" + channel.name + "</span><svg viewBox=\"0 0 1000 1000\"><circle cx=\"500\" cy=\"500\" r=\"500\"></circle><line x1=\"250\" y1=\"750\" x2=\"750\" y2=\"250\"></line><line x1=\"250\" y1=\"250\" x2=\"750\" y2=\"750\"></line></svg>";
	
	// Append tab to the tab bar right before the new tab button
	document.querySelector(".tabs").insertBefore(tabElement, document.querySelector(".tabs .add"));
	
	// Add a new object to the tab list containing the network address and channel name
	tabs.list.push({
		address: network.address,
		channel: channel.name
	});
};

// Prompts the user for information on creating a new tab
var newTab = function() {
	alert("Adding tab");
};

// Returns the network given a tab's index
var getNetworkFromTabIndex = function(tabIndex) {
	// Find the index of the tab's network address in the list of networks
	var networkIndex = networks.map(function(n) {
		return n.address;
	}).indexOf(tabs.list[tabIndex].address);
	
	// Return the network with the found index
	return networks[networkIndex];
};

// Returns a channel given a tab's index
var getChannelFromTabIndex = function(tabIndex) {
	// Get the network
	var network = getNetworkFromTabIndex(tabIndex);
	
	// Find the index of the tab's channel in the network's list of channels
	var channelIndex = network.channels.map(function(c) {
		return c.name;
	}).indexOf(tabs.list[tabIndex].channel);
	
	// Return the channel with the found index
	return network.channels[channelIndex];
};
