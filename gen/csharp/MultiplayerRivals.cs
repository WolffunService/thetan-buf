// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/multiplayer/v1/multiplayer_rivals.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Multiplayer.V1 {

  /// <summary>Holder for reflection information generated from thetan/multiplayer/v1/multiplayer_rivals.proto</summary>
  public static partial class MultiplayerRivalsReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/multiplayer/v1/multiplayer_rivals.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MultiplayerRivalsReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ci50aGV0YW4vbXVsdGlwbGF5ZXIvdjEvbXVsdGlwbGF5ZXJfcml2YWxzLnBy",
            "b3RvEhV0aGV0YW4ubXVsdGlwbGF5ZXIudjEaIXRoZXRhbi9zaGFyZWQvdjEv",
            "Y3VzdG9taXplZC5wcm90byKBBAoVUml2YWxzUGxheWVySW5mb1Byb3RvEhoK",
            "CG1pbmlvbklkGAEgASgJUghtaW5pb25JZBISCgRza2luGAIgASgFUgRza2lu",
            "EmgKDmNvc21ldGljSW5Vc2VkGAMgAygLMkAudGhldGFuLm11bHRpcGxheWVy",
            "LnYxLlJpdmFsc1BsYXllckluZm9Qcm90by5Db3NtZXRpY0luVXNlZEVudHJ5",
            "Ug5jb3NtZXRpY0luVXNlZBI8CgpjdXN0b21pemVkGAQgASgLMhwudGhldGFu",
            "LnNoYXJlZC52MS5DdXN0b21pemVkUgpjdXN0b21pemVkEmsKD2FkZEluQ3Vz",
            "dG9taXplZBgFIAMoCzJBLnRoZXRhbi5tdWx0aXBsYXllci52MS5SaXZhbHNQ",
            "bGF5ZXJJbmZvUHJvdG8uQWRkSW5DdXN0b21pemVkRW50cnlSD2FkZEluQ3Vz",
            "dG9taXplZBpBChNDb3NtZXRpY0luVXNlZEVudHJ5EhAKA2tleRgBIAEoCVID",
            "a2V5EhQKBXZhbHVlGAIgASgDUgV2YWx1ZToCOAEaYAoUQWRkSW5DdXN0b21p",
            "emVkRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSMgoFdmFsdWUYAiABKAsyHC50",
            "aGV0YW4uc2hhcmVkLnYxLkN1c3RvbWl6ZWRSBXZhbHVlOgI4AULoAQoZY29t",
            "LnRoZXRhbi5tdWx0aXBsYXllci52MUIWTXVsdGlwbGF5ZXJSaXZhbHNQcm90",
            "b1ABWj10aGV0YW4tYnVmL2dlbi9nby90aGV0YW4vbXVsdGlwbGF5ZXIvdjE7",
            "dGhldGFuX211bHRpcGxheWVyX3YxogIDVE1YqgIVVGhldGFuLk11bHRpcGxh",
            "eWVyLlYxygIVVGhldGFuXE11bHRpcGxheWVyXFYx4gIhVGhldGFuXE11bHRp",
            "cGxheWVyXFYxXEdQQk1ldGFkYXRh6gIXVGhldGFuOjpNdWx0aXBsYXllcjo6",
            "VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Thetan.Shared.V1.CustomizedReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Multiplayer.V1.RivalsPlayerInfoProto), global::Thetan.Multiplayer.V1.RivalsPlayerInfoProto.Parser, new[]{ "MinionId", "Skin", "CosmeticInUsed", "Customized", "AddInCustomized" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, null, })
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class RivalsPlayerInfoProto : pb::IMessage<RivalsPlayerInfoProto>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<RivalsPlayerInfoProto> _parser = new pb::MessageParser<RivalsPlayerInfoProto>(() => new RivalsPlayerInfoProto());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<RivalsPlayerInfoProto> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Multiplayer.V1.MultiplayerRivalsReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RivalsPlayerInfoProto() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RivalsPlayerInfoProto(RivalsPlayerInfoProto other) : this() {
      minionId_ = other.minionId_;
      skin_ = other.skin_;
      cosmeticInUsed_ = other.cosmeticInUsed_.Clone();
      customized_ = other.customized_ != null ? other.customized_.Clone() : null;
      addInCustomized_ = other.addInCustomized_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RivalsPlayerInfoProto Clone() {
      return new RivalsPlayerInfoProto(this);
    }

    /// <summary>Field number for the "minionId" field.</summary>
    public const int MinionIdFieldNumber = 1;
    private string minionId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string MinionId {
      get { return minionId_; }
      set {
        minionId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "skin" field.</summary>
    public const int SkinFieldNumber = 2;
    private int skin_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Skin {
      get { return skin_; }
      set {
        skin_ = value;
      }
    }

    /// <summary>Field number for the "cosmeticInUsed" field.</summary>
    public const int CosmeticInUsedFieldNumber = 3;
    private static readonly pbc::MapField<string, long>.Codec _map_cosmeticInUsed_codec
        = new pbc::MapField<string, long>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForInt64(16, 0L), 26);
    private readonly pbc::MapField<string, long> cosmeticInUsed_ = new pbc::MapField<string, long>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, long> CosmeticInUsed {
      get { return cosmeticInUsed_; }
    }

    /// <summary>Field number for the "customized" field.</summary>
    public const int CustomizedFieldNumber = 4;
    private global::Thetan.Shared.V1.Customized customized_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Shared.V1.Customized Customized {
      get { return customized_; }
      set {
        customized_ = value;
      }
    }

    /// <summary>Field number for the "addInCustomized" field.</summary>
    public const int AddInCustomizedFieldNumber = 5;
    private static readonly pbc::MapField<string, global::Thetan.Shared.V1.Customized>.Codec _map_addInCustomized_codec
        = new pbc::MapField<string, global::Thetan.Shared.V1.Customized>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForMessage(18, global::Thetan.Shared.V1.Customized.Parser), 42);
    private readonly pbc::MapField<string, global::Thetan.Shared.V1.Customized> addInCustomized_ = new pbc::MapField<string, global::Thetan.Shared.V1.Customized>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, global::Thetan.Shared.V1.Customized> AddInCustomized {
      get { return addInCustomized_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as RivalsPlayerInfoProto);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(RivalsPlayerInfoProto other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (MinionId != other.MinionId) return false;
      if (Skin != other.Skin) return false;
      if (!CosmeticInUsed.Equals(other.CosmeticInUsed)) return false;
      if (!object.Equals(Customized, other.Customized)) return false;
      if (!AddInCustomized.Equals(other.AddInCustomized)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (MinionId.Length != 0) hash ^= MinionId.GetHashCode();
      if (Skin != 0) hash ^= Skin.GetHashCode();
      hash ^= CosmeticInUsed.GetHashCode();
      if (customized_ != null) hash ^= Customized.GetHashCode();
      hash ^= AddInCustomized.GetHashCode();
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
      if (MinionId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(MinionId);
      }
      if (Skin != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Skin);
      }
      cosmeticInUsed_.WriteTo(output, _map_cosmeticInUsed_codec);
      if (customized_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Customized);
      }
      addInCustomized_.WriteTo(output, _map_addInCustomized_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (MinionId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(MinionId);
      }
      if (Skin != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Skin);
      }
      cosmeticInUsed_.WriteTo(ref output, _map_cosmeticInUsed_codec);
      if (customized_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Customized);
      }
      addInCustomized_.WriteTo(ref output, _map_addInCustomized_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (MinionId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(MinionId);
      }
      if (Skin != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Skin);
      }
      size += cosmeticInUsed_.CalculateSize(_map_cosmeticInUsed_codec);
      if (customized_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Customized);
      }
      size += addInCustomized_.CalculateSize(_map_addInCustomized_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(RivalsPlayerInfoProto other) {
      if (other == null) {
        return;
      }
      if (other.MinionId.Length != 0) {
        MinionId = other.MinionId;
      }
      if (other.Skin != 0) {
        Skin = other.Skin;
      }
      cosmeticInUsed_.Add(other.cosmeticInUsed_);
      if (other.customized_ != null) {
        if (customized_ == null) {
          Customized = new global::Thetan.Shared.V1.Customized();
        }
        Customized.MergeFrom(other.Customized);
      }
      addInCustomized_.Add(other.addInCustomized_);
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
            MinionId = input.ReadString();
            break;
          }
          case 16: {
            Skin = input.ReadInt32();
            break;
          }
          case 26: {
            cosmeticInUsed_.AddEntriesFrom(input, _map_cosmeticInUsed_codec);
            break;
          }
          case 34: {
            if (customized_ == null) {
              Customized = new global::Thetan.Shared.V1.Customized();
            }
            input.ReadMessage(Customized);
            break;
          }
          case 42: {
            addInCustomized_.AddEntriesFrom(input, _map_addInCustomized_codec);
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
            MinionId = input.ReadString();
            break;
          }
          case 16: {
            Skin = input.ReadInt32();
            break;
          }
          case 26: {
            cosmeticInUsed_.AddEntriesFrom(ref input, _map_cosmeticInUsed_codec);
            break;
          }
          case 34: {
            if (customized_ == null) {
              Customized = new global::Thetan.Shared.V1.Customized();
            }
            input.ReadMessage(Customized);
            break;
          }
          case 42: {
            addInCustomized_.AddEntriesFrom(ref input, _map_addInCustomized_codec);
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
