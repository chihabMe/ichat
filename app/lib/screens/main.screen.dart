import 'package:flutter/material.dart';
import 'package:ichat/screens/auth/signin.screen.dart';
import 'package:ichat/screens/home/home.screen.dart';
import 'package:shared_preferences/shared_preferences.dart';

class MainScreen extends StatefulWidget {
  @override
  _MainScreenState createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  bool _isAuth = false;
  @override
  void initState() {
    super.initState();
    this.checkLoginStatus();
  }

  Future<void> checkLoginStatus() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    bool loggedIn = prefs.getBool('isLoggedIn') ?? false;
    setState(() {
      _isAuth = loggedIn;
    });
  }

  void updateLoginStatus(bool status) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setBool("isAuth", status);
    setState(() {
      _isAuth = status;
    });
  }

  @override
  Widget build(BuildContext context) {
    return _isAuth
        ? HomeScreen(updateLoginStatus: updateLoginStatus)
        : SigninScreen(
            updateLoginStatus: updateLoginStatus,
          );
  }
}
