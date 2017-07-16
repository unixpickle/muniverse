(function() {

  var LOAD_TIMEOUT = 60000;

  window.indexedDB.deleteDatabase('localforage');
  localStorage.clear();

  // Stub out the com4j API to make Construct2 work.
  window.MGEvent = function(name) {
    this.name = name;
  }

  MGEvent.ENTER_GAME = "ENTER_GAME";
  MGEvent.LOAD_COMPLETE = "LOAD_COMPLETE";
  MGEvent.CLICK_CREDITS = "CLICK_CREDITS";
  MGEvent.CLICK_MINILOGO = "CLICK_MINILOGO";
  MGEvent.CLICK_MORE = "CLICK_MORE";
  MGEvent.DOWNLOAD_APP = "DOWNLOAD_APP";
  MGEvent.PAUSE_GAME = "PAUSE_GAME";
  MGEvent.SCREENSHOT = "SCREENSHOT";
  MGEvent.SHARE = "SHARE";
  MGEvent.SHOW_LOSE = "SHOW_LOSE";
  MGEvent.SHOW_WIN = "SHOW_WIN";
  MGEvent.LEVEL_FAIL = "LEVEL_FAIL";
  MGEvent.LEVEL_WIN = "LEVEL_WIN";
  MGEvent.START_GAME = "START_GAME";
  MGEvent.CONTINUE_GAME = "CONTINUE_GAME";
  MGEvent.ADDED_TO_STAGE = "ADDED_TO_STAGE";
  MGEvent.CHANGE_SCENE = "CHANGE_SCENE";
  MGEvent.FULLSCREEN = "FULLSCREEN";

  class MGDelegate {
    constructor() {
      this._listeners = {};
    }

    addEventListener(name, handler, self) {
      if (!this._listeners[name]) {
        this._listeners[name] = [];
      }
      this._listeners[name].push({handler: handler, self: self});
    }

    dispatcherEvent(evt) {
      if (this._listeners[evt.name]) {
        this._listeners[evt.name].forEach((listener) => {
          listener.handler.call(listener.self);
        });
      }
    }
  }

  window.MGDelegate = new MGDelegate;

  window.com4j = {
    config: {
      ForJoyH5_ShowMoreGamesButton: false,
      ForJoyH5_ShowCreditsButton: false,
      ForJoyH5_Fullscreen: false,
      ForJoyH5_SmallLogo: 'images/morebtn-sheet0.png',
      canvasName: 'gameCanvas'
    },
    loading: {
      hideLoadingGif: () => false
    }
  }

  window.init4j = function() {
    $(document.body).append('<canvas id="gameCanvas"></canvas>');
    cr_createRuntime(com4j.config.canvasName);
    window.cr_getC2Runtime = function() {
      var canvas = document.getElementById("gameCanvas");
      if (canvas) {
        return canvas.c2runtime;
      } else if (window.c2runtime) {
        return window.c2runtime;
      } else {
        return null;
      }
    };

    return pollAndWait(LOAD_TIMEOUT, function() {
      return window.MGDelegate._listeners[MGEvent.ENTER_GAME];
    }).then(function() {
      window.MGDelegate.dispatcherEvent(new MGEvent(MGEvent.ENTER_GAME));
      faketime.advance(50);
    });
  };

})();
