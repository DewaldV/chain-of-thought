function chainsController($scope) {
	$scope.name = "chainsController";

	$scope.aValue = "Some magic value!";

	$scope.button = function() {
		$scope.aValue = "Some other magic value!"
	};
}