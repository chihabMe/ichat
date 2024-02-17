import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'dart:async';
import 'package:dio/dio.dart';
import 'package:ichat/utils/api.constants.dart';

class AuthService {
  final storage = const FlutterSecureStorage();
  final dio = Dio();
  AuthService() {
    _setupInterceptors();
  }
  void _setupInterceptors() {}

  Future<bool> authenticate(String email, String password) async {
    try {
      final response = await dio.post(ApiConstants.OBTAIN_TOKEN_ENDPOINT,
          data: {'email': email, 'password': 'password'});
      if (response.statusCode == 200) {
        final data = response.data;
        final accessToken = data['access_token'];
        final refreshToken = data['refresh_token'];
        await storage.write(key: "access_token", value: accessToken);
        await storage.write(key: "refresh_token", value: refreshToken);
        return true;
      } else {
        return false;
      }
    } catch (e) {
      return false;
    }
  }

  Future<void> refreshAccessToken() async {
    final refreshToken = await storage.read(key: "refresh_token");
    if (refreshToken != null) {
      try {
        final response = await dio.post(ApiConstants.REFRESH_TOKEN_ENDPOINT,
            data: {'refresh_token': refreshToken});
        final accessToken = response.data['access_token'];
        if (accessToken != null) {
          await storage.write(key: "access_token", value: accessToken);
        }
        return;
      } catch (e) {
        print("failed to refresh access token $e");
      }
    } else {
      this.logout();
    }
  }

  Future<void> logout() async {
    await storage.delete(key: "refresh_token");
    await storage.delete(key: "access_token");
  }
}
