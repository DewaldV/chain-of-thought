function leadsController($scope,leadsBind)
{
	$scope.email = "";
	$scope.testClick = function() {
		$scope.Lead = leadsBind.get( { email: $scope.email  } );
	}	
	$scope.button = function() {
		leadsBind.save( { email: $scope.email },
			function(success,status,header)
			{
				$scope.Response = 'THANKS!';
			},
			function(errors)
			{
				$scope.Response = 'Ooooooh Something Bad Happend...';
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

