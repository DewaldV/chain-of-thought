function leadsController($scope)
{
	$scope.inEmail = "";
	$scope.email = "";
	$scope.button = function() {
		$scope.email = $scope.inEmail;
		alert($scope.email);
	};

}

function chainsController($scope) {
	$scope.name = "chainsController";

	$scope.aValue = "Some magic value!";

	$scope.button = function() {
		$scope.aValue = "Some other magic value!"
	};
}

