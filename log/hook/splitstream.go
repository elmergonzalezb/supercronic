package hook

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type writerHook struct {
	writer io.Writer
	levels []logrus.Level
}

func (h *writerHook) Levels() []logrus.Level {
	return h.levels
}

func (h *writerHook) Fire(entry *logrus.Entry) error {
	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(serialized)
	return err
}

func RegisterSplitLogger(logger *logrus.Logger, outWriter io.Writer, errWriter io.Writer) {
	logger.SetOutput(ioutil.Discard)

	logger.AddHook(&writerHook{
		writer: outWriter,
		levels: []logrus.Level{
			logrus.DebugLevel,
			logrus.InfoLevel,
		},
	})

	logger.AddHook(&writerHook{
		writer: errWriter,
		levels: []logrus.Level{
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	})
}
