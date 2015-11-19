
package gui

type connectionSettingsDialogDef  struct{}

func (w connectionSettingsDialogDef) getDefinition() string {
	return `
<interface>
  <object class="GtkDialog" id="ConnectionSettingsDialog">
    <property name="default-width">300</property>
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="title" translatable="yes">Could not connect to account</property>
    <child internal-child="vbox">
      <object class="GtkBox" id="box">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object class="GtkLabel" id="label1">
            <property name="label" translatable="yes">Server (for google accounts, you may want to try xmpp.google.com)</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="server">
            <property name="has-focus">true</property>
            <signal name="activate" handler="reconnect" />
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="label2">
            <property name="label" translatable="yes">Port</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="port">
            <signal name="activate" handler="reconnect" />
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">3</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="save">
            <property name="label" translatable="yes">Connect</property>
            <signal name="clicked" handler="reconnect"/>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">4</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>

`
}
