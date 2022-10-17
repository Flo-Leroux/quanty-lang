// Generated from /mnt/projects/quanty/pkg/grammar/quanty.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class quantyParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, IDENT=4, STRING=5, NUMBER=6, WS=7;
	public static final int
		RULE_schema = 0, RULE_component = 1, RULE_selections = 2, RULE_selection = 3,
		RULE_field = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"schema", "component", "selections", "selection", "field"
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

	@Override
	public String getGrammarFileName() { return "quanty.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public quantyParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SchemaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(quantyParser.EOF, 0); }
		public List<ComponentContext> component() {
			return getRuleContexts(ComponentContext.class);
		}
		public ComponentContext component(int i) {
			return getRuleContext(ComponentContext.class,i);
		}
		public SchemaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_schema; }
	}

	public final SchemaContext schema() throws RecognitionException {
		SchemaContext _localctx = new SchemaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_schema);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(11);
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(10);
				component();
				}
				}
				setState(13);
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__0 );
			setState(15);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComponentContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(quantyParser.IDENT, 0); }
		public SelectionsContext selections() {
			return getRuleContext(SelectionsContext.class,0);
		}
		public ComponentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_component; }
	}

	public final ComponentContext component() throws RecognitionException {
		ComponentContext _localctx = new ComponentContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_component);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(17);
			match(T__0);
			setState(18);
			match(IDENT);
			setState(19);
			match(T__1);
			setState(20);
			selections();
			setState(21);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectionsContext extends ParserRuleContext {
		public List<SelectionContext> selection() {
			return getRuleContexts(SelectionContext.class);
		}
		public SelectionContext selection(int i) {
			return getRuleContext(SelectionContext.class,i);
		}
		public SelectionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selections; }
	}

	public final SelectionsContext selections() throws RecognitionException {
		SelectionsContext _localctx = new SelectionsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_selections);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			selection();
			setState(27);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT || _la==STRING) {
				{
				{
				setState(24);
				selection();
				}
				}
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectionContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(quantyParser.STRING, 0); }
		public FieldContext field() {
			return getRuleContext(FieldContext.class,0);
		}
		public SelectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selection; }
	}

	public final SelectionContext selection() throws RecognitionException {
		SelectionContext _localctx = new SelectionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_selection);
		try {
			setState(32);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(30);
				match(STRING);
				}
				break;
			case IDENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				field();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(quantyParser.IDENT, 0); }
		public SelectionsContext selections() {
			return getRuleContext(SelectionsContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_field);
		try {
			setState(40);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(34);
				match(IDENT);
				setState(35);
				match(T__1);
				setState(36);
				selections();
				setState(37);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				match(IDENT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\t-\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\6\2\16\n\2\r\2\16\2\17\3\2\3\2\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\4\3\4\7\4\34\n\4\f\4\16\4\37\13\4\3\5\3\5\5\5#\n\5"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\5\6+\n\6\3\6\2\2\7\2\4\6\b\n\2\2\2+\2\r\3\2\2"+
		"\2\4\23\3\2\2\2\6\31\3\2\2\2\b\"\3\2\2\2\n*\3\2\2\2\f\16\5\4\3\2\r\f\3"+
		"\2\2\2\16\17\3\2\2\2\17\r\3\2\2\2\17\20\3\2\2\2\20\21\3\2\2\2\21\22\7"+
		"\2\2\3\22\3\3\2\2\2\23\24\7\3\2\2\24\25\7\6\2\2\25\26\7\4\2\2\26\27\5"+
		"\6\4\2\27\30\7\5\2\2\30\5\3\2\2\2\31\35\5\b\5\2\32\34\5\b\5\2\33\32\3"+
		"\2\2\2\34\37\3\2\2\2\35\33\3\2\2\2\35\36\3\2\2\2\36\7\3\2\2\2\37\35\3"+
		"\2\2\2 #\7\7\2\2!#\5\n\6\2\" \3\2\2\2\"!\3\2\2\2#\t\3\2\2\2$%\7\6\2\2"+
		"%&\7\4\2\2&\'\5\6\4\2\'(\7\5\2\2(+\3\2\2\2)+\7\6\2\2*$\3\2\2\2*)\3\2\2"+
		"\2+\13\3\2\2\2\6\17\35\"*";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
