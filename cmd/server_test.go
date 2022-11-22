package cmd_test

import (
	"context"
	"testing"

	cmd "github.com/LordCeilan/go-gopher-grpc/cmd"
	pb "github.com/LordCeilan/go-gopher-grpc/pkg/gopher"
	. "github.com/onsi/gomega"
)

func TestGetGopher(t *testing.T) {
	s := cmd.Server{}

	testCases := []struct {
		name        string
		req         *pb.GopherRequest
		message     string
		expectedErr bool
	}{
		{
			name:        "req ok",
			req:         &pb.GopherRequest{Name: "yoda-gopher"},
			message:     "https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png\n",
			expectedErr: false,
		},
		{
			name:        "req with empty name",
			req:         &pb.GopherRequest{},
			expectedErr: true,
		},
		{
			name:        "nil request",
			req:         nil,
			expectedErr: true,
		},
		{
			name: "Incorrect request",
			req:  &pb.GopherRequest{Name: "y"},
			message: `https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png
https://raw.githubusercontent.com/scraly/gophers/main/LICENSE
https://raw.githubusercontent.com/scraly/gophers/main/arrow-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png
https://raw.githubusercontent.com/scraly/gophers/main/baywatch.png
https://raw.githubusercontent.com/scraly/gophers/main/bike-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/blues-gophers.png
https://raw.githubusercontent.com/scraly/gophers/main/buffy-the-gopher-slayer.png
https://raw.githubusercontent.com/scraly/gophers/main/chandleur-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/cherry-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/cloud-nord.png
https://raw.githubusercontent.com/scraly/gophers/main/devnation-france-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/dr-who.png
https://raw.githubusercontent.com/scraly/gophers/main/dumbledore.png
https://raw.githubusercontent.com/scraly/gophers/main/fire-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/firefly-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/fort-boyard.png
https://raw.githubusercontent.com/scraly/gophers/main/friends.png
https://raw.githubusercontent.com/scraly/gophers/main/gandalf.png
https://raw.githubusercontent.com/scraly/gophers/main/gladiator-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/gopher-dead.png
https://raw.githubusercontent.com/scraly/gophers/main/gopher-johnny.jpg
https://raw.githubusercontent.com/scraly/gophers/main/gopher-open.png
https://raw.githubusercontent.com/scraly/gophers/main/gopher-speaker.png
https://raw.githubusercontent.com/scraly/gophers/main/gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/graffiti-devfest-nantes-2021.png
https://raw.githubusercontent.com/scraly/gophers/main/halloween-spider.png
https://raw.githubusercontent.com/scraly/gophers/main/happy-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/harry-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/idea-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/indiana-jones.png
https://raw.githubusercontent.com/scraly/gophers/main/jedi-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/jurassic-park.png
https://raw.githubusercontent.com/scraly/gophers/main/love-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/luigi-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/mac-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/marshal.png
https://raw.githubusercontent.com/scraly/gophers/main/men-in-black-v2.png
https://raw.githubusercontent.com/scraly/gophers/main/mojito-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/paris-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/pere-fouras.png
https://raw.githubusercontent.com/scraly/gophers/main/recipe.png
https://raw.githubusercontent.com/scraly/gophers/main/sandcastle-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/santa-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/saved-by-the-bell.png
https://raw.githubusercontent.com/scraly/gophers/main/sheldon.png
https://raw.githubusercontent.com/scraly/gophers/main/star-wars.png
https://raw.githubusercontent.com/scraly/gophers/main/stargate.png
https://raw.githubusercontent.com/scraly/gophers/main/tadx-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/unicorn.png
https://raw.githubusercontent.com/scraly/gophers/main/urgences.png
https://raw.githubusercontent.com/scraly/gophers/main/vampire-xmas.png
https://raw.githubusercontent.com/scraly/gophers/main/wired-gopher.png
https://raw.githubusercontent.com/scraly/gophers/main/x-files.png
https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png
`,
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := context.Background()

			response, err := s.GetGopher(ctx, testCase.req)
			t.Log("Got : ", response)

			if testCase.expectedErr {
				g.Expect(response).ToNot(BeNil(), "Result should be nil")
				g.Expect(err).ToNot(BeNil(), "Result should be nil")

			} else {
				g.Expect(response.Message).To(Equal(testCase.message))
			}

		})
	}
}
