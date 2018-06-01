import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'api/api.dart';
import 'api/talkaneast.dart';
import 'widget_message_input.dart';
import 'package:rxdart/rxdart.dart';
import 'models/message.dart';

void main() => runApp(new MyApp());

class MyApp extends StatelessWidget {
  final ChatApi chatApi = new TalkaneastApi();

  @override
  Widget build(BuildContext context) {
    chatApi.connect();
    final title = 'Talkaneast';

    return new MaterialApp(
      title: title,
      home: new MyHomePage(
        chatApi: chatApi,
        title: title,
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  final ChatApi chatApi;
  final String title;

  MyHomePage({Key key, @required this.title, @required this.chatApi})
      : super(key: key);

  @override
  _MyHomePageState createState() => new _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List<Message> messages = new List();

  @override
  void initState() {
    super.initState();
    widget.chatApi.messageSubject.stream.listen(_handleMessage);
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: new AppBar(
        title: new Text(widget.title),
      ),
      body: new Padding(
        padding: const EdgeInsets.all(20.0),
        child: new Column(
          children: <Widget>[
            new Expanded(
              child: new ListView(
              shrinkWrap: true,
              reverse: true,
              children: messages.map((Message message) {
                return new Row(children: 
                [ new Text(message.username + " :" + message.message)]
                );
              }).toList(),
            ),
            ),
            new MessageInputWidget(
              chatApi: widget.chatApi,
            )
          ],
        ),
      ),
    );
  }

  void _handleMessage(Message message) {
    setState(() {
      messages.insert(0, message);
    });
  }

  @override
  void dispose() {
    super.dispose();
    widget.chatApi.dispose();
  }
}