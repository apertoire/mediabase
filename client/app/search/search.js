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
        vm.setWatched = setWatched;

        console.log('activated search view');
        $scope.$onRootScope('/local/search', doSearch);

        function doSearch(me, term) {
            console.log('searching for me: '+me+'term: '+term);
            return api.searchMovies(term).then(function(data) {
                vm.movies = data;
                return vm.movies;
            })
        };

        function setWatched(idx) {
            console.log("maldecido!!!!: ", idx);
            var index = idx;
            return api.setWatched(vm.movies[idx]).then(function(data) {
                console.log('renacuajo!!!: ', vm.movies[index]);
            })
        };

        // function changeDate(movie, watchedDate) {
        //     console.log("yist: ", watchedDate);

        // };

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