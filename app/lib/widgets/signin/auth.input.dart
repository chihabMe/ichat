import 'package:flutter/material.dart';

class AuthInput extends StatefulWidget {
  late TextEditingController controller;
  late String label;
  AuthInput({required this.controller, required this.label});
  @override
  _AuthInputState createState() =>
      _AuthInputState(controller: this.controller, label: this.label);
}

class _AuthInputState extends State<AuthInput> {
  TextEditingController controller;
  late String label;
  _AuthInputState({required this.controller, required this.label});
  @override
  Widget build(BuildContext context) {
    return Container(
        child: Column(
      children: [
        TextField(
          controller: this.controller,
          decoration: InputDecoration(label: Text(this.label)),
        )
      ],
    ));
  }
}
