import 'package:flutter/material.dart';

class SocialSignin extends StatelessWidget {
  const SocialSignin({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Row(mainAxisAlignment: MainAxisAlignment.center, children: [
        SocialSigninItem(
            imagePath: "assets/social_icons/facebook.png", onPress: () {}),
        SocialSigninItem(
            imagePath: "assets/social_icons/google.png", onPress: () {}),
        SocialSigninItem(
            imagePath: "assets/social_icons/apple.png", onPress: () {}),
      ]),
    );
  }
}

class SocialSigninItem extends StatelessWidget {
  late String imagePath;
  late Function onPress;
  SocialSigninItem({required this.imagePath, required this.onPress}) {}

  @override
  Widget build(BuildContext context) {
    return TextButton(
        onPressed: () {
          this.onPress();
        },
        child: Container(
            margin: EdgeInsets.all(5),
            padding: EdgeInsets.all(12),
            decoration: BoxDecoration(
                border: Border.all(color: Colors.black, width: 1),
                borderRadius: BorderRadius.circular(1000)),
            child: Image(
              image: AssetImage(
                this.imagePath,
              ),
              width: 32,
              height: 32,
            )));
  }
}
