(function () {
    'use strict';

    angular
        .module('app.search')
        .controller('Search', Search);

    // Search.$inject = ['$q', 'api', 'logger'];

    /* @ngInject */
    function Search($q, $scope, api, logger) {

        /*jshint validthis: true */
        var vm = this;

        vm.movies = [];
        vm.saveMovie = saveMovie;

        console.log('activated search view');
        $scope.$onRootScope('/local/search', doSearch);

        function doSearch(me, term) {
            console.log('searching for me: '+me+'term: '+term);
            return api.searchMovies(term).then(function(data) {
                vm.movies = data;
                return vm.movies;
            })
        }

        function saveMovie(idx) {
            console.log("maldecido!!!!: ", vm.movies[idx]);
            var index = idx;
            return api.saveMovie(vm.movies[idx]).then(function(data) {
                vm.movies[index] = data;
            })
        }

//         function activate() {
//             return getRecentMovies().then(function() {
//                 logger.info('activated recent view');
//             });
// //             var promises = [getAvengerCount(), getAvengersCast()];
// // //            Using a resolver on all routes or dataservice.ready in every controller
// // //            return dataservice.ready(promises).then(function(){
// //             return $q.all(promises).then(function(){
// //                 logger.info('Activated Dashboard View');
// //             });
//         }

//         function getRecentMovies() {
//             return api.getRecentMovies().then(function (data) {
//                 logger.info('what is: ', data)
//                 vm.movies = data;
//                 return vm.movies;
//             });
//         }
    }
})();