import 'package:flutter/material.dart';
import 'package:ichat/screens/auth/signup.screen.dart';
import 'package:ichat/screens/home/home.screen.dart';
import 'package:ichat/widgets/signin/login.form.dart';
import 'package:ichat/widgets/ui/limited.text.dart';

class SigninScreen extends StatelessWidget {
  final Function(bool) updateLoginStatus;

  SigninScreen({required this.updateLoginStatus});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: SingleChildScrollView(
          scrollDirection: Axis.vertical,
          child: Container(
            padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
            width: double.infinity,
            child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  SizedBox(height: 130),
                  Text("Login in to IChat",
                      style:
                          TextStyle(fontSize: 25, fontWeight: FontWeight.bold)),
                  SizedBox(height: 15),
                  Container(
                    width: 270,
                    child: Text(
                        "Welcome back to ichat please enter your credentials to login into your account",
                        style: TextStyle(color: Colors.black87)),
                  ),
                  LoginForm(),
                ]),
          ),
        ),
      ),
    );
  }
}
