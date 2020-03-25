// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// BlockBlobOperations contains the methods for the BlockBlob group.
type BlockBlobOperations interface {
	// CommitBlockList - The Commit Block List operation writes a blob by specifying the list of block IDs that make up the blob. In order to be written as part of a blob, a block must have been successfully written to the server in a prior Put Block operation. You can call Put Block List to update a blob by uploading only those blocks that have changed, then committing the new and existing blocks together. You can do this by specifying whether to commit a block from the committed block list or from the uncommitted block list, or to commit the most recently uploaded version of the block, whichever list it may belong to.
	CommitBlockList(ctx context.Context, blocks BlockLookupList, options *BlockBlobCommitBlockListOptions) (*BlockBlobCommitBlockListResponse, error)
	// GetBlockList - The Get Block List operation retrieves the list of blocks that have been uploaded as part of a block blob
	GetBlockList(ctx context.Context, listType BlockListType, options *BlockBlobGetBlockListOptions) (*BlockListResponse, error)
	// StageBlock - The Stage Block operation creates a new block to be committed as part of a blob
	StageBlock(ctx context.Context, blockId string, contentLength int64, options *BlockBlobStageBlockOptions) (*BlockBlobStageBlockResponse, error)
	// StageBlockFromURL - The Stage Block operation creates a new block to be committed as part of a blob where the contents are read from a URL.
	StageBlockFromURL(ctx context.Context, blockId string, contentLength int64, sourceUrl url.URL, options *BlockBlobStageBlockFromURLOptions) (*BlockBlobStageBlockFromURLResponse, error)
	// Upload - The Upload Block Blob operation updates the content of an existing block blob. Updating an existing block blob overwrites any existing metadata on the blob. Partial updates are not supported with Put Blob; the content of the existing blob is overwritten with the content of the new blob. To perform a partial update of the content of a block blob, use the Put Block List operation.
	Upload(ctx context.Context, contentLength int64, options *BlockBlobUploadOptions) (*BlockBlobUploadResponse, error)
}

// blockBlobOperations implements the BlockBlobOperations interface.
type blockBlobOperations struct {
	*Client
	urlParameter string
}

// CommitBlockList - The Commit Block List operation writes a blob by specifying the list of block IDs that make up the blob. In order to be written as part of a blob, a block must have been successfully written to the server in a prior Put Block operation. You can call Put Block List to update a blob by uploading only those blocks that have changed, then committing the new and existing blocks together. You can do this by specifying whether to commit a block from the committed block list or from the uncommitted block list, or to commit the most recently uploaded version of the block, whichever list it may belong to.
func (client *blockBlobOperations) CommitBlockList(ctx context.Context, blocks BlockLookupList, options *BlockBlobCommitBlockListOptions) (*BlockBlobCommitBlockListResponse, error) {
	req, err := client.commitBlockListCreateRequest(client.urlParameter, blocks, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.commitBlockListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// commitBlockListCreateRequest creates the CommitBlockList request.
func (client *blockBlobOperations) commitBlockListCreateRequest(urlParameter string, blocks BlockLookupList, options *BlockBlobCommitBlockListOptions) (*azcore.Request, error) {
	urlPath := "/{containerName}/{blob}"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("comp", "blocklist")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	if options != nil && options.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *options.BlobCacheControl)
	}
	if options != nil && options.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *options.BlobContentType)
	}
	if options != nil && options.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *options.BlobContentEncoding)
	}
	if options != nil && options.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *options.BlobContentLanguage)
	}
	if options != nil && options.BlobContentMd5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(*options.BlobContentMd5))
	}
	if options != nil && options.TransactionalContentMd5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*options.TransactionalContentMd5))
	}
	if options != nil && options.TransactionalContentCrc64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(*options.TransactionalContentCrc64))
	}
	if options != nil && options.Metadata != nil {
		req.Header.Set("x-ms-meta", *options.Metadata)
	}
	if options != nil && options.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *options.LeaseId)
	}
	if options != nil && options.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *options.BlobContentDisposition)
	}
	if options != nil && options.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *options.EncryptionKey)
	}
	if options != nil && options.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *options.EncryptionKeySha256)
	}
	if options != nil && options.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *options.EncryptionScope)
	}
	if options != nil && options.Tier != nil {
		req.Header.Set("x-ms-access-tier", string(*options.Tier))
	}
	if options != nil && options.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", options.IfModifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", options.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	err = req.MarshalAsXML(blocks)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// commitBlockListHandleResponse handles the CommitBlockList response.
func (client *blockBlobOperations) commitBlockListHandleResponse(resp *azcore.Response) (*BlockBlobCommitBlockListResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, newStorageError(resp)
	}
	result := BlockBlobCommitBlockListResponse{RawResponse: resp.Response}
	eTag := resp.Header.Get("ETag")
	result.ETag = &eTag
	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return nil, err
	}
	result.LastModified = &lastModified
	contentMd5, err := base64.StdEncoding.DecodeString(resp.Header.Get("Content-MD5"))
	if err != nil {
		return nil, err
	}
	result.ContentMd5 = &contentMd5
	contentCrc64, err := base64.StdEncoding.DecodeString(resp.Header.Get("x-ms-content-crc64"))
	if err != nil {
		return nil, err
	}
	result.ContentCrc64 = &contentCrc64
	clientRequestId := resp.Header.Get("x-ms-client-request-id")
	result.ClientRequestId = &clientRequestId
	requestId := resp.Header.Get("x-ms-request-id")
	result.RequestId = &requestId
	version := resp.Header.Get("x-ms-version")
	result.Version = &version
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	result.Date = &date
	requestServerEncrypted, err := strconv.ParseBool(resp.Header.Get("x-ms-request-server-encrypted"))
	if err != nil {
		return nil, err
	}
	result.RequestServerEncrypted = &requestServerEncrypted
	encryptionKeySha256 := resp.Header.Get("x-ms-encryption-key-sha256")
	result.EncryptionKeySha256 = &encryptionKeySha256
	encryptionScope := resp.Header.Get("x-ms-encryption-scope")
	result.EncryptionScope = &encryptionScope
	return &result, nil
}

// GetBlockList - The Get Block List operation retrieves the list of blocks that have been uploaded as part of a block blob
func (client *blockBlobOperations) GetBlockList(ctx context.Context, listType BlockListType, options *BlockBlobGetBlockListOptions) (*BlockListResponse, error) {
	req, err := client.getBlockListCreateRequest(client.urlParameter, listType, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBlockListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBlockListCreateRequest creates the GetBlockList request.
func (client *blockBlobOperations) getBlockListCreateRequest(urlParameter string, listType BlockListType, options *BlockBlobGetBlockListOptions) (*azcore.Request, error) {
	urlPath := "/{containerName}/{blob}"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("comp", "blocklist")
	if options != nil && options.Snapshot != nil {
		query.Set("snapshot", *options.Snapshot)
	}
	query.Set("blocklisttype", string(listType))
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	if options != nil && options.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *options.LeaseId)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	return req, nil
}

// getBlockListHandleResponse handles the GetBlockList response.
func (client *blockBlobOperations) getBlockListHandleResponse(resp *azcore.Response) (*BlockListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := BlockListResponse{RawResponse: resp.Response}
	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return nil, err
	}
	result.LastModified = &lastModified
	eTag := resp.Header.Get("ETag")
	result.ETag = &eTag
	contentType := resp.Header.Get("Content-Type")
	result.ContentType = &contentType
	blobContentLength, err := strconv.ParseInt(resp.Header.Get("x-ms-blob-content-length"), 10, 64)
	if err != nil {
		return nil, err
	}
	result.BlobContentLength = &blobContentLength
	clientRequestId := resp.Header.Get("x-ms-client-request-id")
	result.ClientRequestId = &clientRequestId
	requestId := resp.Header.Get("x-ms-request-id")
	result.RequestId = &requestId
	version := resp.Header.Get("x-ms-version")
	result.Version = &version
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	result.Date = &date
	return &result, resp.UnmarshalAsXML(&result.BlockList)
}

// StageBlock - The Stage Block operation creates a new block to be committed as part of a blob
func (client *blockBlobOperations) StageBlock(ctx context.Context, blockId string, contentLength int64, options *BlockBlobStageBlockOptions) (*BlockBlobStageBlockResponse, error) {
	req, err := client.stageBlockCreateRequest(client.urlParameter, blockId, contentLength, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.stageBlockHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// stageBlockCreateRequest creates the StageBlock request.
func (client *blockBlobOperations) stageBlockCreateRequest(urlParameter string, blockId string, contentLength int64, options *BlockBlobStageBlockOptions) (*azcore.Request, error) {
	urlPath := "/{containerName}/{blob}"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("comp", "block")
	query.Set("blockid", blockId)
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if options != nil && options.TransactionalContentMd5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*options.TransactionalContentMd5))
	}
	if options != nil && options.TransactionalContentCrc64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(*options.TransactionalContentCrc64))
	}
	if options != nil && options.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *options.LeaseId)
	}
	if options != nil && options.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *options.EncryptionKey)
	}
	if options != nil && options.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *options.EncryptionKeySha256)
	}
	if options != nil && options.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *options.EncryptionScope)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	return req, nil
}

// stageBlockHandleResponse handles the StageBlock response.
func (client *blockBlobOperations) stageBlockHandleResponse(resp *azcore.Response) (*BlockBlobStageBlockResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, newStorageError(resp)
	}
	result := BlockBlobStageBlockResponse{RawResponse: resp.Response}
	contentMd5, err := base64.StdEncoding.DecodeString(resp.Header.Get("Content-MD5"))
	if err != nil {
		return nil, err
	}
	result.ContentMd5 = &contentMd5
	clientRequestId := resp.Header.Get("x-ms-client-request-id")
	result.ClientRequestId = &clientRequestId
	requestId := resp.Header.Get("x-ms-request-id")
	result.RequestId = &requestId
	version := resp.Header.Get("x-ms-version")
	result.Version = &version
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	result.Date = &date
	contentCrc64, err := base64.StdEncoding.DecodeString(resp.Header.Get("x-ms-content-crc64"))
	if err != nil {
		return nil, err
	}
	result.ContentCrc64 = &contentCrc64
	requestServerEncrypted, err := strconv.ParseBool(resp.Header.Get("x-ms-request-server-encrypted"))
	if err != nil {
		return nil, err
	}
	result.RequestServerEncrypted = &requestServerEncrypted
	encryptionKeySha256 := resp.Header.Get("x-ms-encryption-key-sha256")
	result.EncryptionKeySha256 = &encryptionKeySha256
	encryptionScope := resp.Header.Get("x-ms-encryption-scope")
	result.EncryptionScope = &encryptionScope
	return &result, nil
}

// StageBlockFromURL - The Stage Block operation creates a new block to be committed as part of a blob where the contents are read from a URL.
func (client *blockBlobOperations) StageBlockFromURL(ctx context.Context, blockId string, contentLength int64, sourceUrl url.URL, options *BlockBlobStageBlockFromURLOptions) (*BlockBlobStageBlockFromURLResponse, error) {
	req, err := client.stageBlockFromUrlCreateRequest(client.urlParameter, blockId, contentLength, sourceUrl, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.stageBlockFromUrlHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// stageBlockFromUrlCreateRequest creates the StageBlockFromURL request.
func (client *blockBlobOperations) stageBlockFromUrlCreateRequest(urlParameter string, blockId string, contentLength int64, sourceUrl url.URL, options *BlockBlobStageBlockFromURLOptions) (*azcore.Request, error) {
	urlPath := "/{containerName}/{blob}"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("comp", "block")
	query.Set("blockid", blockId)
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("x-ms-copy-source", sourceUrl.String())
	if options != nil && options.SourceRange != nil {
		req.Header.Set("x-ms-source-range", *options.SourceRange)
	}
	if options != nil && options.SourceContentMd5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(*options.SourceContentMd5))
	}
	if options != nil && options.SourceContentcrc64 != nil {
		req.Header.Set("x-ms-source-content-crc64", base64.StdEncoding.EncodeToString(*options.SourceContentcrc64))
	}
	if options != nil && options.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *options.EncryptionKey)
	}
	if options != nil && options.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *options.EncryptionKeySha256)
	}
	if options != nil && options.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *options.EncryptionScope)
	}
	if options != nil && options.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *options.LeaseId)
	}
	if options != nil && options.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", options.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", options.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *options.SourceIfMatch)
	}
	if options != nil && options.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *options.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	return req, nil
}

// stageBlockFromUrlHandleResponse handles the StageBlockFromURL response.
func (client *blockBlobOperations) stageBlockFromUrlHandleResponse(resp *azcore.Response) (*BlockBlobStageBlockFromURLResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, newStorageError(resp)
	}
	result := BlockBlobStageBlockFromURLResponse{RawResponse: resp.Response}
	contentMd5, err := base64.StdEncoding.DecodeString(resp.Header.Get("Content-MD5"))
	if err != nil {
		return nil, err
	}
	result.ContentMd5 = &contentMd5
	clientRequestId := resp.Header.Get("x-ms-client-request-id")
	result.ClientRequestId = &clientRequestId
	requestId := resp.Header.Get("x-ms-request-id")
	result.RequestId = &requestId
	version := resp.Header.Get("x-ms-version")
	result.Version = &version
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	result.Date = &date
	contentCrc64, err := base64.StdEncoding.DecodeString(resp.Header.Get("x-ms-content-crc64"))
	if err != nil {
		return nil, err
	}
	result.ContentCrc64 = &contentCrc64
	requestServerEncrypted, err := strconv.ParseBool(resp.Header.Get("x-ms-request-server-encrypted"))
	if err != nil {
		return nil, err
	}
	result.RequestServerEncrypted = &requestServerEncrypted
	encryptionKeySha256 := resp.Header.Get("x-ms-encryption-key-sha256")
	result.EncryptionKeySha256 = &encryptionKeySha256
	encryptionScope := resp.Header.Get("x-ms-encryption-scope")
	result.EncryptionScope = &encryptionScope
	return &result, nil
}

// Upload - The Upload Block Blob operation updates the content of an existing block blob. Updating an existing block blob overwrites any existing metadata on the blob. Partial updates are not supported with Put Blob; the content of the existing blob is overwritten with the content of the new blob. To perform a partial update of the content of a block blob, use the Put Block List operation.
func (client *blockBlobOperations) Upload(ctx context.Context, contentLength int64, options *BlockBlobUploadOptions) (*BlockBlobUploadResponse, error) {
	req, err := client.uploadCreateRequest(client.urlParameter, contentLength, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.uploadHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// uploadCreateRequest creates the Upload request.
func (client *blockBlobOperations) uploadCreateRequest(urlParameter string, contentLength int64, options *BlockBlobUploadOptions) (*azcore.Request, error) {
	urlPath := "/{containerName}/{blob}"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header.Set("x-ms-blob-type", "BlockBlob")
	if options != nil && options.TransactionalContentMd5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*options.TransactionalContentMd5))
	}
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if options != nil && options.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *options.BlobContentType)
	}
	if options != nil && options.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *options.BlobContentEncoding)
	}
	if options != nil && options.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *options.BlobContentLanguage)
	}
	if options != nil && options.BlobContentMd5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(*options.BlobContentMd5))
	}
	if options != nil && options.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *options.BlobCacheControl)
	}
	if options != nil && options.Metadata != nil {
		req.Header.Set("x-ms-meta", *options.Metadata)
	}
	if options != nil && options.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *options.LeaseId)
	}
	if options != nil && options.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *options.BlobContentDisposition)
	}
	if options != nil && options.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *options.EncryptionKey)
	}
	if options != nil && options.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *options.EncryptionKeySha256)
	}
	if options != nil && options.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *options.EncryptionScope)
	}
	if options != nil && options.Tier != nil {
		req.Header.Set("x-ms-access-tier", string(*options.Tier))
	}
	if options != nil && options.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", options.IfModifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", options.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	return req, nil
}

// uploadHandleResponse handles the Upload response.
func (client *blockBlobOperations) uploadHandleResponse(resp *azcore.Response) (*BlockBlobUploadResponse, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, newStorageError(resp)
	}
	result := BlockBlobUploadResponse{RawResponse: resp.Response}
	eTag := resp.Header.Get("ETag")
	result.ETag = &eTag
	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return nil, err
	}
	result.LastModified = &lastModified
	contentMd5, err := base64.StdEncoding.DecodeString(resp.Header.Get("Content-MD5"))
	if err != nil {
		return nil, err
	}
	result.ContentMd5 = &contentMd5
	clientRequestId := resp.Header.Get("x-ms-client-request-id")
	result.ClientRequestId = &clientRequestId
	requestId := resp.Header.Get("x-ms-request-id")
	result.RequestId = &requestId
	version := resp.Header.Get("x-ms-version")
	result.Version = &version
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	result.Date = &date
	requestServerEncrypted, err := strconv.ParseBool(resp.Header.Get("x-ms-request-server-encrypted"))
	if err != nil {
		return nil, err
	}
	result.RequestServerEncrypted = &requestServerEncrypted
	encryptionKeySha256 := resp.Header.Get("x-ms-encryption-key-sha256")
	result.EncryptionKeySha256 = &encryptionKeySha256
	encryptionScope := resp.Header.Get("x-ms-encryption-scope")
	result.EncryptionScope = &encryptionScope
	return &result, nil
}