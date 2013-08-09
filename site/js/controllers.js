function leadsController($scope,leadsBind)
{
	$scope.email = "";
	$scope.button = function() {
		alert($scope.email)
		leadsBind.save($scope.email,
			function(success,status,header)
			{
				alert('SAVED!');
			},
			function(errors)
			{
				alert('SHIT IS FUCKED!');
			});
	};

}

function chainsController($scope) {
	$scope.name = "chainsController";

	$scope.aValue = "Some magic value!";

	$scope.button = function() {
		$scope.aValue = "Some other magic value!"
	};
}

