cb.init = function() {
	var run = function() {
		// Add click listener to tab buttons
		document.querySelector(".tabs").addEventListener("click", cb.tabs.tabClickEvents);
		
		// Initially open a tab for each channel in each network
		for (var network in cb.connection.networks) {
			for (var channel in cb.connection.networks[network].channels) {
				cb.tabs.addTab(cb.connection.networks[network], cb.connection.networks[network].channels[channel]);
			}
		}
		
		// Initially load the default tab
		cb.tabs.changeTab(cb.tabs.tabs.active);
	};
	
	return {
		run: run
	};
}();
