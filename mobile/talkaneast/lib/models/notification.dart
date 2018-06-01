class NotificationMessage {
  final String message;

  NotificationMessage(this.message);

  NotificationMessage.fromJson(Map<String, dynamic> json)
    : message = json["message"];
}