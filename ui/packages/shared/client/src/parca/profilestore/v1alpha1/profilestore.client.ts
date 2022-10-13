// @generated by protobuf-ts 2.8.1 with parameter long_type_string,generate_dependencies
// @generated from protobuf file "parca/profilestore/v1alpha1/profilestore.proto" (package "parca.profilestore.v1alpha1", syntax proto3)
// tslint:disable
import { AgentsService } from "./profilestore";
import type { AgentsResponse } from "./profilestore";
import type { AgentsRequest } from "./profilestore";
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { ProfileStoreService } from "./profilestore";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { WriteRawResponse } from "./profilestore";
import type { WriteRawRequest } from "./profilestore";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * ProfileStoreService is the service the accepts pprof writes
 *
 * @generated from protobuf service parca.profilestore.v1alpha1.ProfileStoreService
 */
export interface IProfileStoreServiceClient {
    /**
     * WriteRaw accepts a raw set of bytes of a pprof file
     *
     * @generated from protobuf rpc: WriteRaw(parca.profilestore.v1alpha1.WriteRawRequest) returns (parca.profilestore.v1alpha1.WriteRawResponse);
     */
    writeRaw(input: WriteRawRequest, options?: RpcOptions): UnaryCall<WriteRawRequest, WriteRawResponse>;
}
/**
 * ProfileStoreService is the service the accepts pprof writes
 *
 * @generated from protobuf service parca.profilestore.v1alpha1.ProfileStoreService
 */
export class ProfileStoreServiceClient implements IProfileStoreServiceClient, ServiceInfo {
    typeName = ProfileStoreService.typeName;
    methods = ProfileStoreService.methods;
    options = ProfileStoreService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * WriteRaw accepts a raw set of bytes of a pprof file
     *
     * @generated from protobuf rpc: WriteRaw(parca.profilestore.v1alpha1.WriteRawRequest) returns (parca.profilestore.v1alpha1.WriteRawResponse);
     */
    writeRaw(input: WriteRawRequest, options?: RpcOptions): UnaryCall<WriteRawRequest, WriteRawResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<WriteRawRequest, WriteRawResponse>("unary", this._transport, method, opt, input);
    }
}
/**
 * AgentsService maintains the agents
 *
 * @generated from protobuf service parca.profilestore.v1alpha1.AgentsService
 */
export interface IAgentsServiceClient {
    /**
     * Agents return the agents that pushed data to the server
     *
     * @generated from protobuf rpc: Agents(parca.profilestore.v1alpha1.AgentsRequest) returns (parca.profilestore.v1alpha1.AgentsResponse);
     */
    agents(input: AgentsRequest, options?: RpcOptions): UnaryCall<AgentsRequest, AgentsResponse>;
}
/**
 * AgentsService maintains the agents
 *
 * @generated from protobuf service parca.profilestore.v1alpha1.AgentsService
 */
export class AgentsServiceClient implements IAgentsServiceClient, ServiceInfo {
    typeName = AgentsService.typeName;
    methods = AgentsService.methods;
    options = AgentsService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Agents return the agents that pushed data to the server
     *
     * @generated from protobuf rpc: Agents(parca.profilestore.v1alpha1.AgentsRequest) returns (parca.profilestore.v1alpha1.AgentsResponse);
     */
    agents(input: AgentsRequest, options?: RpcOptions): UnaryCall<AgentsRequest, AgentsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<AgentsRequest, AgentsResponse>("unary", this._transport, method, opt, input);
    }
}
