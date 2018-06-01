class Message {
  final String username;
  final String message;

  Message(this.username, this.message);

  Message.fromJson(Map<String, dynamic> json)
    : username = json["username"],
    message = json["message"];

    Map<String, dynamic> toJson() =>
    {
      'event' : 'Message',
      'data': {
        'message': this.message,
      }
    };
}