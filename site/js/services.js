angular.module('leadsService', ['ngResource'])
.factory('leadsBind', function($resource) { 
	return $resource('http://localhost\\:8787/leads/:id', { id: '@id' }); 
});
