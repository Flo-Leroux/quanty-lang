// Generated from /mnt/projects/quanty/pkg/grammar/quanty.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class quantyLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, IDENT=4, STRING=5, NUMBER=6, WS=7;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "IDENT", "STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT",
			"NUMBER", "INT", "EXP", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'component'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "IDENT", "STRING", "NUMBER", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public quantyLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "quanty.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\tq\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\3\3\3\3\4\3\4\3\5\3\5\7\5.\n\5\f\5\16\5\61\13\5\3\6\3\6\3\6\7\6\66"+
		"\n\6\f\6\16\69\13\6\3\6\3\6\3\7\3\7\3\7\5\7@\n\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\t\3\t\3\n\3\n\3\13\5\13M\n\13\3\13\3\13\3\13\6\13R\n\13\r\13\16\13"+
		"S\5\13V\n\13\3\13\5\13Y\n\13\3\f\3\f\3\f\7\f^\n\f\f\f\16\fa\13\f\5\fc"+
		"\n\f\3\r\3\r\5\rg\n\r\3\r\3\r\3\16\6\16l\n\16\r\16\16\16m\3\16\3\16\2"+
		"\2\17\3\3\5\4\7\5\t\6\13\7\r\2\17\2\21\2\23\2\25\b\27\2\31\2\33\t\3\2"+
		"\f\3\2c|\5\2\62;C\\c|\n\2$$\61\61^^ddhhppttvv\5\2\62;CHch\5\2\2!$$^^\3"+
		"\2\62;\3\2\63;\4\2GGgg\4\2--//\5\2\13\f\17\17\"\"\2v\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\25\3\2\2\2\2\33\3\2\2\2"+
		"\3\35\3\2\2\2\5\'\3\2\2\2\7)\3\2\2\2\t+\3\2\2\2\13\62\3\2\2\2\r<\3\2\2"+
		"\2\17A\3\2\2\2\21G\3\2\2\2\23I\3\2\2\2\25L\3\2\2\2\27b\3\2\2\2\31d\3\2"+
		"\2\2\33k\3\2\2\2\35\36\7e\2\2\36\37\7q\2\2\37 \7o\2\2 !\7r\2\2!\"\7q\2"+
		"\2\"#\7p\2\2#$\7g\2\2$%\7p\2\2%&\7v\2\2&\4\3\2\2\2\'(\7}\2\2(\6\3\2\2"+
		"\2)*\7\177\2\2*\b\3\2\2\2+/\t\2\2\2,.\t\3\2\2-,\3\2\2\2.\61\3\2\2\2/-"+
		"\3\2\2\2/\60\3\2\2\2\60\n\3\2\2\2\61/\3\2\2\2\62\67\7$\2\2\63\66\5\r\7"+
		"\2\64\66\5\23\n\2\65\63\3\2\2\2\65\64\3\2\2\2\669\3\2\2\2\67\65\3\2\2"+
		"\2\678\3\2\2\28:\3\2\2\29\67\3\2\2\2:;\7$\2\2;\f\3\2\2\2<?\7^\2\2=@\t"+
		"\4\2\2>@\5\17\b\2?=\3\2\2\2?>\3\2\2\2@\16\3\2\2\2AB\7w\2\2BC\5\21\t\2"+
		"CD\5\21\t\2DE\5\21\t\2EF\5\21\t\2F\20\3\2\2\2GH\t\5\2\2H\22\3\2\2\2IJ"+
		"\n\6\2\2J\24\3\2\2\2KM\7/\2\2LK\3\2\2\2LM\3\2\2\2MN\3\2\2\2NU\5\27\f\2"+
		"OQ\7\60\2\2PR\t\7\2\2QP\3\2\2\2RS\3\2\2\2SQ\3\2\2\2ST\3\2\2\2TV\3\2\2"+
		"\2UO\3\2\2\2UV\3\2\2\2VX\3\2\2\2WY\5\31\r\2XW\3\2\2\2XY\3\2\2\2Y\26\3"+
		"\2\2\2Zc\7\62\2\2[_\t\b\2\2\\^\t\7\2\2]\\\3\2\2\2^a\3\2\2\2_]\3\2\2\2"+
		"_`\3\2\2\2`c\3\2\2\2a_\3\2\2\2bZ\3\2\2\2b[\3\2\2\2c\30\3\2\2\2df\t\t\2"+
		"\2eg\t\n\2\2fe\3\2\2\2fg\3\2\2\2gh\3\2\2\2hi\5\27\f\2i\32\3\2\2\2jl\t"+
		"\13\2\2kj\3\2\2\2lm\3\2\2\2mk\3\2\2\2mn\3\2\2\2no\3\2\2\2op\b\16\2\2p"+
		"\34\3\2\2\2\17\2/\65\67?LSUX_bfm\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
