import 'package:flutter/material.dart';
import 'package:ichat/widgets/signin/auth.input.dart';

class LoginForm extends StatefulWidget {
  @override
  _LoginFormState createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  TextEditingController _email = TextEditingController();
  TextEditingController _password = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(children: [
        AuthInput(controller: this._email, label: "Your email"),
        SizedBox(height: 15),
        AuthInput(
          controller: this._password,
          label: "Password",
          secure: true,
        )
      ]),
    );
  }
}
