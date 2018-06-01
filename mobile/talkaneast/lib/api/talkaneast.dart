import 'api.dart';
import '../models/message.dart';
import '../models/notification.dart';
import 'package:rxdart/rxdart.dart';
import 'dart:io';
import 'dart:convert';


class TalkaneastApi implements ChatApi {
  WebSocket _websocket;

  @override
  PublishSubject<Message> messageSubject = new PublishSubject();

  @override
  PublishSubject<NotificationMessage> notificationSubject = new PublishSubject();

  @override
  void connect() {
    WebSocket.connect("ws://80.211.54.149:2148/ws").then((socket) {
      _websocket = socket;
      _websocket.handleError(_handleError);
      _websocket.listen(_handleSocket);
      changeNick("adam12");
    });
    
  }

  _handleSocket(msg) {
    Map event = json.decode(msg);
    try {
      print(event["event"]);
      switch (event["event"]) {
        case "Message":
          Message message = Message.fromJson(event["data"]);
          messageSubject.add(message);
          break;

        case "Notification":
          NotificationMessage message = NotificationMessage.fromJson(event["data"]);
          notificationSubject.add(message);
          break;
        }
    } catch(e) {}
  }

  _handleError(error) {}

  @override
  void dispose() {
    _websocket.close();
  }

  @override
  void changeNick(String nick) {
    _websocket.add(
      json.encode(
        {
          'event': 'SetNick',
          'data': {
            'nickname': nick,
          }
        }
      )
    );
  }

  @override
  void sendMessage(Message message) {
    _websocket.add(json.encode(message.toJson()));
  }
}