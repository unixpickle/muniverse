(function() {

  localStorage.clear();

  var apiObj = {
    services: {
      connect: () => true,
      addEventListener: () => true,
      getUserId: () => 0,
      getUsername: () => 'Guest'
    },
    stats: {
      submit: (a, b) => console.log('stats.submit('+a+','+b+')')
    }
  };

  window.kongregateAPI = {
    loadAPI: function(cb) {
      setTimeout(cb, 1);
    },
    getAPI: function() {
      return apiObj;
    }
  };

})();
