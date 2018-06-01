import 'package:rxdart/rxdart.dart';
import '../models/message.dart';
import '../models/notification.dart';

abstract class ChatApi {
  PublishSubject<Message> messageSubject;
  PublishSubject<NotificationMessage> notificationSubject;

  void connect();
  void dispose();
  void sendMessage(Message message);
  void changeNick(String nick);
}