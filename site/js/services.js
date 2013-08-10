angular.module('leadsService', ['ngResource']).
	factory('leadsBind', function($resource) { 
		return $resource('http://localhost\\:8787/leads/:email', { email: '@email' }, 
			{
				save: { method: "POST", params: { email: '' } }
			}); 
	});
