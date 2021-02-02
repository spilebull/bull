# logger

- zap.SugaredLogger(Logger から取得できる簡易な Logger。構造化ロギングと printf スタイルのロギングどちらも利用可能ですが、Logger よりも低速で高アロケーションです。)
  これは簡単で一般的な処理を優先使う。
  使用例：
  logger.Info(args ...interface{})
  logger.Infof(template string, args ...interface{})
- zap.Logger(SugaredLogger よりも高速かつ低アロケーションで動作しますが、構造化スタイルのロギングしかできません（print や printf スタイルのロギングはできません)
  使用例：
  logger.Logger.Info(msg string, fields ...zapcore.Field)
