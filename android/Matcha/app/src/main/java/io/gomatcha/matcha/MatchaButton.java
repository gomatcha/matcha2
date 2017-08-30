package io.gomatcha.matcha;

import android.content.Context;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RelativeLayout;

import com.google.protobuf.InvalidProtocolBufferException;

import io.gomatcha.matcha.pb.view.PbView;
import io.gomatcha.matcha.pb.view.alert.PbAlert;
import io.gomatcha.matcha.pb.view.button.PbButton;

public class MatchaButton extends MatchaChildView {
    Button view;

    static {
        MatchaView.registerView("gomatcha.io/matcha/view/button", new MatchaView.ViewFactory() {
            @Override
            public MatchaChildView createView(Context context, MatchaViewNode node) {
                return new MatchaTextView(context, node);
            }
        });
    }

    public MatchaButton(Context context, MatchaViewNode node) {
        super(context, node);

        RelativeLayout.LayoutParams params = new RelativeLayout.LayoutParams(RelativeLayout.LayoutParams.MATCH_PARENT, RelativeLayout.LayoutParams.MATCH_PARENT);
        view = new Button(context);
        addView(view, params);
    }

    @Override
    public void setNode(PbView.BuildNode buildNode) {
        super.setNode(buildNode);
        try {
            PbButton.View proto = buildNode.getBridgeValue().unpack(PbButton.View.class);
        } catch (InvalidProtocolBufferException e) {
        }
    }
}