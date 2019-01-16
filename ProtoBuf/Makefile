release:
		protoc -I . landingcrew.proto3 code_workflow.proto3 data_workflow.proto3 content_workflow.proto3 action_workflow.proto3 --go_out=plugins=grpc:../landingcrew/workflow/
		protoc -I . landingcrew.proto3 code_workflow.proto3 data_workflow.proto3 content_workflow.proto3 action_workflow.proto3 --go_out=plugins=grpc:../api.landingcrew.com/workflow/