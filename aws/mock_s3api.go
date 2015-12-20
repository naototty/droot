package aws

import "github.com/stretchr/testify/mock"

import "github.com/aws/aws-sdk-go/aws/request"
import "github.com/aws/aws-sdk-go/service/s3"

type S3API struct {
	mock.Mock
}

func (_m *S3API) AbortMultipartUploadRequest(_a0 *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.AbortMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.AbortMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.AbortMultipartUploadInput) *s3.AbortMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.AbortMultipartUploadOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) AbortMultipartUpload(_a0 *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.AbortMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.AbortMultipartUploadInput) *s3.AbortMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.AbortMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.AbortMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) CompleteMultipartUploadRequest(_a0 *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CompleteMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CompleteMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.CompleteMultipartUploadInput) *s3.CompleteMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CompleteMultipartUploadOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) CompleteMultipartUpload(_a0 *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CompleteMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.CompleteMultipartUploadInput) *s3.CompleteMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CompleteMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) CopyObjectRequest(_a0 *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CopyObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CopyObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.CopyObjectInput) *s3.CopyObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CopyObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) CopyObject(_a0 *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CopyObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.CopyObjectInput) *s3.CopyObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CopyObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CopyObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) CreateBucketRequest(_a0 *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CreateBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CreateBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.CreateBucketInput) *s3.CreateBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CreateBucketOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) CreateBucket(_a0 *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CreateBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.CreateBucketInput) *s3.CreateBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CreateBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) CreateMultipartUploadRequest(_a0 *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CreateMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CreateMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.CreateMultipartUploadInput) *s3.CreateMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CreateMultipartUploadOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) CreateMultipartUpload(_a0 *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CreateMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.CreateMultipartUploadInput) *s3.CreateMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CreateMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketRequest(_a0 *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInput) *s3.DeleteBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucket(_a0 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInput) *s3.DeleteBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketCorsRequest(_a0 *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketCorsInput) *s3.DeleteBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketCorsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketCors(_a0 *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketCorsInput) *s3.DeleteBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketLifecycleRequest(_a0 *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketLifecycleInput) *s3.DeleteBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketLifecycleOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketLifecycle(_a0 *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketLifecycleInput) *s3.DeleteBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketPolicyRequest(_a0 *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketPolicyInput) *s3.DeleteBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketPolicyOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketPolicy(_a0 *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketPolicyInput) *s3.DeleteBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketReplicationRequest(_a0 *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketReplicationInput) *s3.DeleteBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketReplicationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketReplication(_a0 *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketReplicationInput) *s3.DeleteBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketTaggingRequest(_a0 *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketTaggingInput) *s3.DeleteBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketTaggingOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketTagging(_a0 *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketTaggingInput) *s3.DeleteBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketWebsiteRequest(_a0 *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketWebsiteInput) *s3.DeleteBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketWebsiteOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteBucketWebsite(_a0 *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketWebsiteInput) *s3.DeleteBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteObjectRequest(_a0 *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectInput) *s3.DeleteObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteObject(_a0 *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectInput) *s3.DeleteObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) DeleteObjectsRequest(_a0 *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteObjectsOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectsInput) *s3.DeleteObjectsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteObjectsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) DeleteObjects(_a0 *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteObjectsOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectsInput) *s3.DeleteObjectsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketAclRequest(_a0 *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketAclOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAclInput) *s3.GetBucketAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketAclOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketAcl(_a0 *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketAclOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAclInput) *s3.GetBucketAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketCorsRequest(_a0 *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketCorsInput) *s3.GetBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketCorsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketCors(_a0 *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketCorsInput) *s3.GetBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketLifecycleRequest(_a0 *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleInput) *s3.GetBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLifecycleOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketLifecycle(_a0 *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleInput) *s3.GetBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketLifecycleConfigurationRequest(_a0 *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleConfigurationInput) *s3.GetBucketLifecycleConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLifecycleConfigurationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketLifecycleConfiguration(_a0 *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleConfigurationInput) *s3.GetBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketLocationRequest(_a0 *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLocationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLocationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLocationInput) *s3.GetBucketLocationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLocationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketLocation(_a0 *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLocationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLocationInput) *s3.GetBucketLocationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLocationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLocationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketLoggingRequest(_a0 *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLoggingOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLoggingInput) *s3.GetBucketLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLoggingOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketLogging(_a0 *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLoggingInput) *s3.GetBucketLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketNotificationRequest(_a0 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.NotificationConfigurationDeprecated
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfigurationDeprecated); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.NotificationConfigurationDeprecated)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketNotification(_a0 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	ret := _m.Called(_a0)

	var r0 *s3.NotificationConfigurationDeprecated
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfigurationDeprecated); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfigurationDeprecated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketNotificationConfigurationRequest(_a0 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.NotificationConfiguration
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfiguration); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.NotificationConfiguration)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketNotificationConfiguration(_a0 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	ret := _m.Called(_a0)

	var r0 *s3.NotificationConfiguration
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfiguration); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketPolicyRequest(_a0 *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketPolicyOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketPolicy(_a0 *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketReplicationRequest(_a0 *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketReplicationInput) *s3.GetBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketReplicationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketReplication(_a0 *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketReplicationInput) *s3.GetBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketRequestPaymentRequest(_a0 *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketRequestPaymentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketRequestPaymentOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketRequestPaymentInput) *s3.GetBucketRequestPaymentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketRequestPaymentOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketRequestPayment(_a0 *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketRequestPaymentInput) *s3.GetBucketRequestPaymentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketRequestPaymentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketTaggingRequest(_a0 *s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketTaggingInput) *s3.GetBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketTaggingOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketTagging(_a0 *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketTaggingInput) *s3.GetBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketVersioningRequest(_a0 *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketVersioningInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketVersioningOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketVersioningInput) *s3.GetBucketVersioningOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketVersioningOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketVersioning(_a0 *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketVersioningInput) *s3.GetBucketVersioningOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketVersioningInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetBucketWebsiteRequest(_a0 *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketWebsiteInput) *s3.GetBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketWebsiteOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetBucketWebsite(_a0 *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketWebsiteInput) *s3.GetBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetObjectRequest(_a0 *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectInput) *s3.GetObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetObject(_a0 *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectInput) *s3.GetObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetObjectAclRequest(_a0 *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectAclOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectAclInput) *s3.GetObjectAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectAclOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetObjectAcl(_a0 *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectAclOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectAclInput) *s3.GetObjectAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) GetObjectTorrentRequest(_a0 *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTorrentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectTorrentOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTorrentInput) *s3.GetObjectTorrentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectTorrentOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) GetObjectTorrent(_a0 *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectTorrentOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTorrentInput) *s3.GetObjectTorrentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectTorrentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTorrentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) HeadBucketRequest(_a0 *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.HeadBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.HeadBucketInput) *s3.HeadBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.HeadBucketOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) HeadBucket(_a0 *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.HeadBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) *s3.HeadBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.HeadBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) HeadObjectRequest(_a0 *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.HeadObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.HeadObjectInput) *s3.HeadObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.HeadObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) HeadObject(_a0 *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.HeadObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) *s3.HeadObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.HeadObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListBucketsRequest(_a0 *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListBucketsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListBucketsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListBucketsInput) *s3.ListBucketsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListBucketsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) ListBuckets(_a0 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListBucketsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListBucketsInput) *s3.ListBucketsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListBucketsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListMultipartUploadsRequest(_a0 *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListMultipartUploadsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListMultipartUploadsInput) *s3.ListMultipartUploadsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListMultipartUploadsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) ListMultipartUploads(_a0 *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListMultipartUploadsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput) *s3.ListMultipartUploadsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListMultipartUploadsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListMultipartUploadsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListMultipartUploadsPages(_a0 *s3.ListMultipartUploadsInput, _a1 func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *S3API) ListObjectVersionsRequest(_a0 *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListObjectVersionsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListObjectVersionsInput) *s3.ListObjectVersionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListObjectVersionsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) ListObjectVersions(_a0 *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListObjectVersionsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput) *s3.ListObjectVersionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListObjectVersionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListObjectVersionsPages(_a0 *s3.ListObjectVersionsInput, _a1 func(*s3.ListObjectVersionsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *S3API) ListObjectsRequest(_a0 *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListObjectsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsInput) *s3.ListObjectsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListObjectsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) ListObjects(_a0 *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListObjectsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput) *s3.ListObjectsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListObjectsPages(_a0 *s3.ListObjectsInput, _a1 func(*s3.ListObjectsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *S3API) ListPartsRequest(_a0 *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListPartsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListPartsInput) *s3.ListPartsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListPartsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) ListParts(_a0 *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListPartsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput) *s3.ListPartsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListPartsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListPartsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) ListPartsPages(_a0 *s3.ListPartsInput, _a1 func(*s3.ListPartsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *S3API) PutBucketAclRequest(_a0 *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketAclOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAclInput) *s3.PutBucketAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketAclOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketAcl(_a0 *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketAclOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAclInput) *s3.PutBucketAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketCorsRequest(_a0 *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketCorsInput) *s3.PutBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketCorsOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketCors(_a0 *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketCorsInput) *s3.PutBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketLifecycleRequest(_a0 *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleInput) *s3.PutBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLifecycleOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketLifecycle(_a0 *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleInput) *s3.PutBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketLifecycleConfigurationRequest(_a0 *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleConfigurationInput) *s3.PutBucketLifecycleConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLifecycleConfigurationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketLifecycleConfiguration(_a0 *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleConfigurationInput) *s3.PutBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketLoggingRequest(_a0 *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLoggingOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLoggingInput) *s3.PutBucketLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLoggingOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketLogging(_a0 *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLoggingInput) *s3.PutBucketLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketNotificationRequest(_a0 *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketNotificationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationInput) *s3.PutBucketNotificationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketNotificationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketNotification(_a0 *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketNotificationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationInput) *s3.PutBucketNotificationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketNotificationConfigurationRequest(_a0 *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketNotificationConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationConfigurationInput) *s3.PutBucketNotificationConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketNotificationConfigurationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketNotificationConfiguration(_a0 *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketNotificationConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationConfigurationInput) *s3.PutBucketNotificationConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketPolicyRequest(_a0 *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketPolicyInput) *s3.PutBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketPolicyOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketPolicy(_a0 *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketPolicyInput) *s3.PutBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketReplicationRequest(_a0 *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketReplicationInput) *s3.PutBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketReplicationOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketReplication(_a0 *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketReplicationInput) *s3.PutBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketRequestPaymentRequest(_a0 *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketRequestPaymentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketRequestPaymentOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketRequestPaymentInput) *s3.PutBucketRequestPaymentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketRequestPaymentOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketRequestPayment(_a0 *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketRequestPaymentInput) *s3.PutBucketRequestPaymentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketRequestPaymentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketTaggingRequest(_a0 *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketTaggingInput) *s3.PutBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketTaggingOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketTagging(_a0 *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketTaggingInput) *s3.PutBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketVersioningRequest(_a0 *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketVersioningInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketVersioningInput) *s3.PutBucketVersioningOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketVersioningOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketVersioning(_a0 *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketVersioningInput) *s3.PutBucketVersioningOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketVersioningInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutBucketWebsiteRequest(_a0 *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketWebsiteInput) *s3.PutBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketWebsiteOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutBucketWebsite(_a0 *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketWebsiteInput) *s3.PutBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutObjectRequest(_a0 *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.PutObjectInput) *s3.PutObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutObject(_a0 *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.PutObjectInput) *s3.PutObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) PutObjectAclRequest(_a0 *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutObjectAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutObjectAclOutput
	if rf, ok := ret.Get(1).(func(*s3.PutObjectAclInput) *s3.PutObjectAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutObjectAclOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) PutObjectAcl(_a0 *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutObjectAclOutput
	if rf, ok := ret.Get(0).(func(*s3.PutObjectAclInput) *s3.PutObjectAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutObjectAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) RestoreObjectRequest(_a0 *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.RestoreObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.RestoreObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.RestoreObjectInput) *s3.RestoreObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.RestoreObjectOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) RestoreObject(_a0 *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.RestoreObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.RestoreObjectInput) *s3.RestoreObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.RestoreObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.RestoreObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) UploadPartRequest(_a0 *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.UploadPartInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.UploadPartOutput
	if rf, ok := ret.Get(1).(func(*s3.UploadPartInput) *s3.UploadPartOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.UploadPartOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) UploadPart(_a0 *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.UploadPartOutput
	if rf, ok := ret.Get(0).(func(*s3.UploadPartInput) *s3.UploadPartOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.UploadPartInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *S3API) UploadPartCopyRequest(_a0 *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.UploadPartCopyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.UploadPartCopyOutput
	if rf, ok := ret.Get(1).(func(*s3.UploadPartCopyInput) *s3.UploadPartCopyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.UploadPartCopyOutput)
		}
	}

	return r0, r1
}
func (_m *S3API) UploadPartCopy(_a0 *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.UploadPartCopyOutput
	if rf, ok := ret.Get(0).(func(*s3.UploadPartCopyInput) *s3.UploadPartCopyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartCopyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.UploadPartCopyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
