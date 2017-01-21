namespace go rpc

service RemoteCmd {
  // Current UTC time
  string utc(),

  // Current CPU usage
  string cpu(),

  // Available RAM
  string ram(),

  // Download url into specific folder
  oneway void download(1:string url, 2:string folder),

  // Make computer "say" something
  oneway void say(1:string phrase),

  // Capture and send a screenshot
  binary screenshot()

  // Trigger webhook at specific time
}
