import 'package:flutter/material.dart';
import 'package:ichat/screens/auth/signin.screen.dart';

class HomeScreen extends StatelessWidget {
  final Function(bool) updateLoginStatus;

  HomeScreen({required this.updateLoginStatus});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Home'),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            updateLoginStatus(false);
            Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) => SigninScreen(
                          updateLoginStatus: updateLoginStatus,
                        )));
          },
          child: Text('Logout'),
        ),
      ),
    );
  }
}
