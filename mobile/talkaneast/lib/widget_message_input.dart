import 'package:flutter/material.dart';
import 'api/api.dart';
import 'models/message.dart';

class MessageInputWidget extends StatefulWidget {
  final ChatApi chatApi;

  MessageInputWidget({Key key, @required this.chatApi})
      : super(key: key);

  @override
  _MessageInputWidgetState createState() => new _MessageInputWidgetState();
}

class _MessageInputWidgetState extends State<MessageInputWidget> {
  TextEditingController _controller = new TextEditingController();

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      widget.chatApi.sendMessage(new Message("", _controller.text));
      _controller.text = "";
    }
  }

  @override
  Widget build(BuildContext context) {
    return new Row(
      children: <Widget>[
        new Expanded(
          child: TextField(
            controller: _controller,
            decoration: new InputDecoration(
              border: InputBorder.none,
              hintText: 'Type message'
            ),
          ),
        ),
        new IconButton(
          icon: new Icon(Icons.send, color: Colors.blue[500]),
          onPressed: _sendMessage,
        )
      ],
    );
  }
}