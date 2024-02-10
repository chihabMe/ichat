import 'package:flutter/material.dart';
import 'package:ichat/screens/auth/signin.screen.dart';
import 'package:ichat/screens/home/home.screen.dart';

class SignupScreen extends StatelessWidget {
  final Function(bool) updateLoginStatus;

  SignupScreen(this.updateLoginStatus);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Sign Up'),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            updateLoginStatus(true);
            Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) => HomeScreen(
                          updateLoginStatus: updateLoginStatus,
                        )));
          },
          child: Text('Sign Up'),
        ),
      ),
    );
  }
}
