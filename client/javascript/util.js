cb.util = function() {
	// Returns the index of an element amongst its siblings in its parent
	var indexInParent = function(node) {
	    var children = node.parentNode.childNodes;
	    var index = 0;
	    for (var i = 0; i < children.length; i++) {
	         if (children[i] === node) {
				 return index;
			 } else if (children[i].nodeType === 1) {
				 index++;
			 }
	    }
	    return -1;
	};
	
	return {
		indexInParent: indexInParent
	};
}();
