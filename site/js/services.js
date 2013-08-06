angular.module('leadsService', ['ngResource'])
.service(function() { 
	return $resource('/leads/{id}', { id: '@id' }); 
});