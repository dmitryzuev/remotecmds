namespace go rpc

service RemoteCmd {
  // Current UTC time
  string utc(),

  // Current CPU usage
  string cpu_usage(),

  // Available RAM
  string available_ram(),

  // CPU usage over last hour
  string cpu_usage_last_hour(),

  // Available RAM over last hour
  string available_ram_last_hour(),

  // Download url into specific folder
  string download_url(1:string url, 2:string folder),

  // Make computer "say" something
  string say(1:string phrase),

  // Capture and send a screenshot
  binary screenshot()

  // Trigger webhook at specific time
}
