import 'package:flutter/material.dart';

class Button extends StatefulWidget {
  late String text;
  late Function onPress;
  late Color color;
  late Color textColor;
  Button(
      {required this.text,
      required this.onPress,
      this.color = Colors.teal,
      this.textColor = Colors.white});

  @override
  _ButtonState createState() => _ButtonState(
      text: this.text,
      onPress: this.onPress,
      color: this.color,
      textColor: this.textColor);
}

class _ButtonState extends State<Button> {
  late String text;
  late Function onPress;
  late Color color;
  late Color textColor;
  _ButtonState(
      {required this.text,
      required this.onPress,
      required this.color,
      required this.textColor});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 53,
      child: TextButton(
        style: ButtonStyle(
          shape: MaterialStateProperty.all(RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(12),
          )),
          backgroundColor: MaterialStateProperty.all(this.color),
        ),
        child: Text(this.text, style: TextStyle(color: this.textColor)),
        onPressed: () {
          this.onPress();
        },
      ),
    );
  }
}
