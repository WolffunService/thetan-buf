// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: wsproto/rivals.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Wsproto {

  /// <summary>Holder for reflection information generated from wsproto/rivals.proto</summary>
  public static partial class RivalsReflection {

    #region Descriptor
    /// <summary>File descriptor for wsproto/rivals.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static RivalsReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChR3c3Byb3RvL3JpdmFscy5wcm90bxIHd3Nwcm90byLmAQoVUml2YWxzUGxh",
            "eWVySW5mb1Byb3RvEhoKCG1pbmlvbklkGAEgASgJUghtaW5pb25JZBISCgRz",
            "a2luGAIgASgFUgRza2luEloKDmNvc21ldGljSW5Vc2VkGAMgAygLMjIud3Nw",
            "cm90by5SaXZhbHNQbGF5ZXJJbmZvUHJvdG8uQ29zbWV0aWNJblVzZWRFbnRy",
            "eVIOY29zbWV0aWNJblVzZWQaQQoTQ29zbWV0aWNJblVzZWRFbnRyeRIQCgNr",
            "ZXkYASABKAlSA2tleRIUCgV2YWx1ZRgCIAEoA1IFdmFsdWU6AjgBQnEKC2Nv",
            "bS53c3Byb3RvQgtSaXZhbHNQcm90b1ABWhl0aGV0YW4tYnVmL2dlbi9nby93",
            "c3Byb3RvogIDV1hYqgIHV3Nwcm90b8oCB1dzcHJvdG/iAhNXc3Byb3RvXEdQ",
            "Qk1ldGFkYXRh6gIHV3Nwcm90b2IGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Wsproto.RivalsPlayerInfoProto), global::Wsproto.RivalsPlayerInfoProto.Parser, new[]{ "MinionId", "Skin", "CosmeticInUsed" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
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
      get { return global::Wsproto.RivalsReflection.Descriptor.MessageTypes[0]; }
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
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (MinionId.Length != 0) hash ^= MinionId.GetHashCode();
      if (Skin != 0) hash ^= Skin.GetHashCode();
      hash ^= CosmeticInUsed.GetHashCode();
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
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
