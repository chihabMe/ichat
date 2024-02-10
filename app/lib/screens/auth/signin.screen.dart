import 'package:flutter/material.dart';
import 'package:ichat/screens/auth/signup.screen.dart';
import 'package:ichat/screens/home/home.screen.dart';
import 'package:ichat/widgets/signin/login.form.dart';
import 'package:ichat/widgets/signin/social.signin.dart';
import 'package:ichat/widgets/ui/input.dart';
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
            padding: EdgeInsets.symmetric(vertical: 10, horizontal: 25),
            width: double.infinity,
            child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  SizedBox(height: 120),
                  Text("Log in in to IChat",
                      style:
                          TextStyle(fontSize: 25, fontWeight: FontWeight.bold)),
                  SizedBox(height: 15),
                  Container(
                    width: 270,
                    child: Text(
                        "Welcome back! Sign in using your social account or email to continue ",
                        textAlign: TextAlign.center,
                        style: TextStyle(
                            color: const Color.fromARGB(255, 134, 134, 134),
                            wordSpacing: 1)),
                  ),
                  SizedBox(height: 15),
                  SocialSignin(),
                  SizedBox(height: 15),
                  LoginForm(),
                  SizedBox(height: 55),
                  Button(text: "login", onPress: () {}),
                  SizedBox(height: 15),
                  Button(
                      text: "Forgot password?",
                      color: Colors.transparent,
                      textColor: Colors.teal,
                      onPress: () {}),
                ]),
          ),
        ),
      ),
    );
  }
}
