import 'package:flutter/material.dart';

class LimitedWordsText extends StatelessWidget {
  final String text;
  final int maxWords;

  LimitedWordsText({required this.text, this.maxWords = 10});

  @override
  Widget build(BuildContext context) {
    List<String> words = text.split(' ');
    String limitedText = words.take(maxWords).join(' ');

    return Text(
      limitedText,
      overflow: TextOverflow.ellipsis,
    );
  }
}
