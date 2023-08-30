// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/gateway/v1/gs_rivalslobby.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Gateway.V1 {

  /// <summary>Holder for reflection information generated from thetan/gateway/v1/gs_rivalslobby.proto</summary>
  public static partial class GsRivalslobbyReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/gateway/v1/gs_rivalslobby.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static GsRivalslobbyReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiZ0aGV0YW4vZ2F0ZXdheS92MS9nc19yaXZhbHNsb2JieS5wcm90bxIRdGhl",
            "dGFuLmdhdGV3YXkudjEaJXRoZXRhbi9yaXZhbHMvdjEvc2VydmljZV9yaXZh",
            "bHMucHJvdG8iZAoSVG93bkFsbG9jYXRpb25SZXNwEhoKCHNlcnZlcklQGAEg",
            "ASgJUghzZXJ2ZXJJUBIeCgpzZXJ2ZXJQb3J0GAIgASgNUgpzZXJ2ZXJQb3J0",
            "EhIKBHRvd24YAyABKAlSBHRvd24ycAoYVGhldGFuR2F0ZXdheVJpdmFsc0xv",
            "YmJ5ElQKDEFsbG9jYXRlVG93bhIbLnRoZXRhbi5yaXZhbHMudjEuTG9iYnlU",
            "b3duGiUudGhldGFuLmdhdGV3YXkudjEuVG93bkFsbG9jYXRpb25SZXNwIgBC",
            "yAEKFWNvbS50aGV0YW4uZ2F0ZXdheS52MUISR3NSaXZhbHNsb2JieVByb3Rv",
            "UAFaNXRoZXRhbi1idWYvZ2VuL2dvL3RoZXRhbi9nYXRld2F5L3YxO3RoZXRh",
            "bl9nYXRld2F5X3YxogIDVEdYqgIRVGhldGFuLkdhdGV3YXkuVjHKAhFUaGV0",
            "YW5cR2F0ZXdheVxWMeICHVRoZXRhblxHYXRld2F5XFYxXEdQQk1ldGFkYXRh",
            "6gITVGhldGFuOjpHYXRld2F5OjpWMWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Thetan.Rivals.V1.ServiceRivalsReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Gateway.V1.TownAllocationResp), global::Thetan.Gateway.V1.TownAllocationResp.Parser, new[]{ "ServerIP", "ServerPort", "Town" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class TownAllocationResp : pb::IMessage<TownAllocationResp>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<TownAllocationResp> _parser = new pb::MessageParser<TownAllocationResp>(() => new TownAllocationResp());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<TownAllocationResp> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Gateway.V1.GsRivalslobbyReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public TownAllocationResp() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public TownAllocationResp(TownAllocationResp other) : this() {
      serverIP_ = other.serverIP_;
      serverPort_ = other.serverPort_;
      town_ = other.town_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public TownAllocationResp Clone() {
      return new TownAllocationResp(this);
    }

    /// <summary>Field number for the "serverIP" field.</summary>
    public const int ServerIPFieldNumber = 1;
    private string serverIP_ = "";
    /// <summary>
    ///fishnet ip+port
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string ServerIP {
      get { return serverIP_; }
      set {
        serverIP_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "serverPort" field.</summary>
    public const int ServerPortFieldNumber = 2;
    private uint serverPort_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public uint ServerPort {
      get { return serverPort_; }
      set {
        serverPort_ = value;
      }
    }

    /// <summary>Field number for the "town" field.</summary>
    public const int TownFieldNumber = 3;
    private string town_ = "";
    /// <summary>
    ///town info
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Town {
      get { return town_; }
      set {
        town_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as TownAllocationResp);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(TownAllocationResp other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (ServerIP != other.ServerIP) return false;
      if (ServerPort != other.ServerPort) return false;
      if (Town != other.Town) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (ServerIP.Length != 0) hash ^= ServerIP.GetHashCode();
      if (ServerPort != 0) hash ^= ServerPort.GetHashCode();
      if (Town.Length != 0) hash ^= Town.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (ServerIP.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ServerIP);
      }
      if (ServerPort != 0) {
        output.WriteRawTag(16);
        output.WriteUInt32(ServerPort);
      }
      if (Town.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Town);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (ServerIP.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(ServerIP);
      }
      if (ServerPort != 0) {
        output.WriteRawTag(16);
        output.WriteUInt32(ServerPort);
      }
      if (Town.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Town);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (ServerIP.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ServerIP);
      }
      if (ServerPort != 0) {
        size += 1 + pb::CodedOutputStream.ComputeUInt32Size(ServerPort);
      }
      if (Town.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Town);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(TownAllocationResp other) {
      if (other == null) {
        return;
      }
      if (other.ServerIP.Length != 0) {
        ServerIP = other.ServerIP;
      }
      if (other.ServerPort != 0) {
        ServerPort = other.ServerPort;
      }
      if (other.Town.Length != 0) {
        Town = other.Town;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            ServerIP = input.ReadString();
            break;
          }
          case 16: {
            ServerPort = input.ReadUInt32();
            break;
          }
          case 26: {
            Town = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            ServerIP = input.ReadString();
            break;
          }
          case 16: {
            ServerPort = input.ReadUInt32();
            break;
          }
          case 26: {
            Town = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code