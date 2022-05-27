package kafka

import "testing"

func Test_namespaceTopicName(t *testing.T) {
	t.Parallel()

	type args struct {
		topicName     string
		envLookupFunc func(string) (string, bool)
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no topic namespacing",
			args: args{
				topicName: unicastMessageV4Topic,
				envLookupFunc: func(s string) (string, bool) {
					return "", false
				},
			},
			want: unicastMessageV4Topic,
		},
		{
			name: "namespace topics",
			args: args{
				topicName: unicastMessageV4Topic,
				envLookupFunc: func(s string) (string, bool) {
					if s == topicNamespaceEnvVariableName {
						return "codebgp", true
					}
					return "", false
				},
			},
			want: "codebgp." + unicastMessageV4Topic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := namespaceTopicName(tt.args.topicName, tt.args.envLookupFunc); got != tt.want {
				t.Errorf("namespaceTopicName() = %v, want %v", got, tt.want)
			}
		})
	}
}
