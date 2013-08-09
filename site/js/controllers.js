function leadsController($scope,leadsBind)
{
	$scope.email = "";
	$scope.button = function() {
		leadsBind.save( { email: $scope.email },
			function(success,status,header)
			{
				$scope.Response = 'SAVED!';
			},
			function(errors)
			{
				$scope.Response = 'SHIT IS FUCKED!';
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

