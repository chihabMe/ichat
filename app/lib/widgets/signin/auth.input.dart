import 'package:flutter/material.dart';

class AuthInput extends StatefulWidget {
  late TextEditingController controller;
  late String label;
  late bool secure;
  AuthInput(
      {required this.controller, required this.label, this.secure = false});
  @override
  _AuthInputState createState() => _AuthInputState(
      controller: this.controller, label: this.label, secure: this.secure);
}

class _AuthInputState extends State<AuthInput> {
  TextEditingController controller;
  late String label;
  late bool secure;
  _AuthInputState(
      {required this.controller, required this.label, required this.secure});
  @override
  Widget build(BuildContext context) {
    return Container(
        child: Column(
      children: [
        TextField(
          controller: this.controller,
          obscureText: this.secure,
          decoration: InputDecoration(
              label: Text(
            this.label,
            style: TextStyle(color: Colors.teal, fontWeight: FontWeight.w500),
          )),
        )
      ],
    ));
  }
}
