// Create a Promise that waits for a condition.
function pollAndWait(timeout, check) {
  var INTERVAL = 100;
  return new Promise(function(resolve, reject) {
    var ival;
    ival = setInterval(function() {
      try {
        if (!check()) {
          timeout -= INTERVAL;
          if (timeout < 0) {
            clearInterval(ival);
            reject('timeout surpassed');
          }
          return;
        }
      } catch (e) {
        reject(e);
        return;
      }
      clearInterval(ival);
      resolve();
    }, INTERVAL);
  });
}
